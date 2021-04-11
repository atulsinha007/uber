package server

import (
	"github.com/atulsinha007/uber/config"
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"github.com/atulsinha007/uber/pkg/log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type serverConfig struct {
	*mux.Router
	Address      string
	serviceName  string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Endpoint struct {
	Path    string
	Method  string
	Handler func(req *http.Request) handler.Response
}

var sc serverConfig

func Init() {
	serviceName := config.V.GetString("SERVICE_NAME")
	port := config.V.GetString("SERVICE_PORT")

	if serviceName == "" {
		log.L.Fatal("Mandatory config SERVICE_NAME missing")
	}

	if port == "" {
		log.L.Fatal("Mandatory config SERVICE_PORT missing")
	}

	readTimeout := config.V.GetInt64("READ_TIMEOUT_MILLIS")
	writeTimeout := config.V.GetInt64("WRITE_TIMEOUT_MILLIS")

	if readTimeout == 0 {
		log.L.Warn("missing READ_TIMEOUT_MILLIS config")
		readTimeout = int64(time.Minute / time.Millisecond)
	}

	if writeTimeout == 0 {
		log.L.Warn("missing WRITE_TIMEOUT_MILLIS config")
		writeTimeout = int64(time.Minute / time.Millisecond)
	}

	router := mux.NewRouter()

	addr := ":" + port
	sc = serverConfig{
		Router:       router,
		Address:      addr,
		serviceName:  serviceName,
		readTimeout:  time.Duration(readTimeout) * time.Millisecond,
		writeTimeout: time.Duration(writeTimeout) * time.Millisecond,
	}
}

func RegisterAndStart(endPts []Endpoint) {
	sc.HandleFunc("/status", handler.StatusActive)

	for _, endPt := range endPts {
		sc.HandleFunc(endPt.Path, handler.Make(endPt.Handler)).Methods(endPt.Method)
	}

	start()
}

func start() {
	headers := handlers.AllowedHeaders([]string{"authorization", "content-type", "x-requested-with"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	srv := &http.Server{
		Handler:      handlers.CORS(headers, origins, methods)(sc.Router),
		Addr:         sc.Address,
		ReadTimeout:  sc.readTimeout,
		WriteTimeout: sc.writeTimeout,
	}

	log.L.With(zap.String("serviceName", sc.serviceName), zap.String("address", sc.Address)).
		Info("server starting")
	log.L.With(zap.Error(srv.ListenAndServe())).Fatal("server error")
}
