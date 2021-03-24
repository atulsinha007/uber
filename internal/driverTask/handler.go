package driverTask

import (
	"encoding/json"
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	ctrl Ctrl
}

func NewHandler(ctrl Ctrl) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) AcceptRideRequest(req *http.Request) handler.Response {
	driverTaskId, ok := mux.Vars(req)["driverTaskId"]
	if !ok || driverTaskId == "" {
		return handler.BadRequest("invalid driverId")
	}

	var payload AcceptRideReq
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	payload.DriverTaskId, _ = strconv.Atoi(driverTaskId)

	err = h.ctrl.AcceptRideRequest(payload)
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
			"data": "ride accepted successfully",
		},
	}
}

func (h *Handler) UpdateRide(req *http.Request) handler.Response {
	driverTaskId, ok := mux.Vars(req)["driverTaskId"]
	if !ok || driverTaskId == "" {
		return handler.BadRequest("invalid driverId")
	}

	var payload UpdateRideReq
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	payload.DriverTaskId, _ = strconv.Atoi(driverTaskId)

	err = h.ctrl.UpdateRide(payload)
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
			"data": "ride updated successfully",
		},
	}
}
