package driverTask

import (
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	ctrl Ctrl
}

func NewHandler(ctrl Ctrl) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetDriverHistory(req *http.Request) handler.Response {
	driverId, ok := mux.Vars(req)["driverId"]
	if !ok || driverId == "" {
		return handler.BadRequest("invalid driverId")
	}

	resp, err := h.ctrl.GetDriverHistory(driverId)
	if err != nil {
		return handler.Response{
			Code: http.StatusInternalServerError,
			Payload: handler.Fields{
				"error": err.Error(),
			},
		}
	}

	return handler.Response{
		Code: http.StatusCreated,
		Payload: handler.Fields{
			"data": resp,
		},
	}
}
