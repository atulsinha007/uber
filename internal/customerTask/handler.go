package customerTask

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

func (h *Handler) CancelRide(req *http.Request) handler.Response {
	customerTaskId, ok := mux.Vars(req)["customerTaskId"]
	if !ok || customerTaskId == "" {
		return handler.BadRequest("invalid customerTaskId")
	}

	id, _ := strconv.Atoi(customerTaskId)

	err := h.ctrl.CancelRide(id)
	if err != nil {
		return handler.Response{
			Code: http.StatusInternalServerError,
			Payload: handler.Fields{
				"error": err.Error(),
			},
		}
	}

	return handler.Response{
		Code: http.StatusOK,
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

	id, _ := strconv.Atoi(customerId)

	resp, err := h.ctrl.GetHistory(id)
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
		Code: http.StatusOK,
		Payload: handler.Fields{
			"data": resp,
		},
	}
}
