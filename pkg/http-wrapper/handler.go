package handler

import (
	"encoding/json"
	"github.com/atulsinha007/uber/pkg/log"
	"go.uber.org/zap"
	"net/http"
)

// Fields represents JSON
type Fields map[string]interface{}

// Response is written to http.ResponseWriter
type Response struct {
	Code    int
	Payload interface{}
}

// Make creates a http handler from a request handler func
func Make(
	f func(req *http.Request) Response,
) func(w http.ResponseWriter, req *http.Request) {
	handler := func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			req.Body.Close()
			return
		}
		w.Header().Set("Content-Type", "application/json")
		res := f(req)
		JSON, err := json.Marshal(res.Payload)
		if err != nil {
			log.L.With(zap.Error(err)).Error("json marshal failed")
		}

		w.WriteHeader(res.Code)
		w.Write(JSON)
		req.Body.Close()

	}

	return handler
}

func StatusActive(w http.ResponseWriter, req *http.Request) {
	Make(func(req *http.Request) Response {
		return Response{
			Code:    http.StatusOK,
			Payload: Fields{"status": "active"},
		}
	})(w, req)
}

func BadRequest(msg string) Response {
	return Response{
		http.StatusBadRequest,
		Fields{"error": msg},
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
