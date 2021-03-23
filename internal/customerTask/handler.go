package customerTask

import (
	"encoding/json"
	"github.com/atulsinha007/uber/internal/address"
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

func (h *Handler) CreateRideRequest(req *http.Request) handler.Response {

	var createRideRequest CreateRideRequest
	err := json.NewDecoder(req.Body).Decode(&createRideRequest)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	resp, err := h.ctrl.CreateRide(createRideRequest)
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

func (h *Handler) UpdateRideStops(req *http.Request) handler.Response {
	customerTaskId, ok := mux.Vars(req)["customerTaskId"]
	if !ok || customerTaskId == "" {
		return handler.BadRequest("invalid customerTaskId")
	}

	var stops []address.Location
	err := json.NewDecoder(req.Body).Decode(&stops)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	updateRideReq := UpdateRideReq{
		CustomerTaskId: customerTaskId,
		Stops:          stops,
	}

	err = h.ctrl.UpdateRide(updateRideReq)
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
			"data": "resp",
		},
	}
}

func (h *Handler) CancelRide(req *http.Request) handler.Response {
	customerTaskId, ok := mux.Vars(req)["customerTaskId"]
	if !ok || customerTaskId == "" {
		return handler.BadRequest("invalid customerTaskId")
	}

	err := h.ctrl.CancelRide(customerTaskId)
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
			"data": "ride cancelled successfully",
		},
	}
}

func (h *Handler) GetCustomerHistory(req *http.Request) handler.Response {
	customerId, ok := mux.Vars(req)["customerId"]
	if !ok || customerId == "" {
		return handler.BadRequest("invalid customerId")
	}

	resp, err := h.ctrl.GetHistory(customerId)
	if err != nil {
		if err.Error() == "record not found" {
			return handler.Response{
				Code:    http.StatusNotFound,
				Payload: handler.Fields{
					"error": err.Error(),
				},
			}
		}
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
