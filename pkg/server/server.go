package server

import (
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	AuthReq bool
}

var sc serverConfig

func init() {
	serviceName := viper.GetString("SERVICE_NAME")
	port := viper.GetString("SERVICE_PORT")

	if serviceName == "" {
		logrus.Fatal("Mandatory config SERVICE_NAME missing")
	}

	if port == "" {
		logrus.Fatal("Mandatory config SERVICE_PORT missing")
	}

	readTimeout := viper.GetInt64("READ_TIMEOUT_MILLIS")
	writeTimeout := viper.GetInt64("WRITE_TIMEOUT_MILLIS")

	if readTimeout == 0 {
		logrus.Warn("missing READ_TIMEOUT_MILLIS config")
		readTimeout = int64(time.Minute / time.Millisecond)
	}

	if writeTimeout == 0 {
		logrus.Warn("missing WRITE_TIMEOUT_MILLIS config")
		writeTimeout = int64(time.Minute / time.Millisecond)
	}

	pathPrefix := "/" + serviceName
	router := mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	addr := ":" + port
	sc = serverConfig{
		Router:       router,
		Address:      addr,
		serviceName:  serviceName,
		readTimeout:  time.Duration(readTimeout) * time.Millisecond,
		writeTimeout: time.Duration(writeTimeout) * time.Millisecond,
	}
}

// Registers handlers for all endpoints and then starts the server
// Adds a basic health-check endpoint at /status. This only checks the connectivity to the serverConfig
// and not actual health e.g. DB connectivity, crucial upstream services availability etc.
func RegisterAndStart(endPts []Endpoint) {
	for _, endPt := range endPts {
		if endPt.AuthReq {
			handlerFunc := http.HandlerFunc(handler.Make(endPt.Handler))
			//newrelic.Instrument(sc.Router, endPt.Path, authentication.ServeHTTP).Methods(endPt.Method)
		} else {
			//newrelic.Instrument(sc.Router, endPt.Path, handler.Make(endPt.Handler)).Methods(endPt.Method)
		}

	}

	sc.Router, "/status", handler.StatusActive).Methods(http.MethodGet)

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

	logrus.Infof("server %v starting at addr: %v", sc.serviceName, sc.Address)
	logrus.Fatal(srv.ListenAndServe())
}
