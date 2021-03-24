package user

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

func (h *Handler) CreateUser(req *http.Request) handler.Response {
	var payload CreateUserRequest
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	err = h.ctrl.AddUser(CreateUserRequestToUser(payload))
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
			"data": "user creation successful",
		},
	}
}

func (h *Handler) GetDriverProfile(req *http.Request) handler.Response {
	driverId, ok := mux.Vars(req)["driverId"]
	if !ok || driverId == "" {
		return handler.BadRequest("invalid driverId")
	}

	id, _ := strconv.Atoi(driverId)

	resp, err := h.ctrl.GetDriverProfile(id)
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
			"data": resp,
		},
	}
}


func (h *Handler) UpdateLocation(req *http.Request) handler.Response {
	userId, ok := mux.Vars(req)["userId"]
	if !ok || userId == "" {
		return handler.BadRequest("invalid userId")
	}

	id, _ := strconv.Atoi(userId)

	var latLng LatLng
	err := json.NewDecoder(req.Body).Decode(&latLng)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	err = h.ctrl.UpdateLocation(UpdateCurrentLocationRequest{
		UserId:      id,
		CurLocation: latLng,
	})
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
			"data": "location updated successfully",
		},
	}
}

func (h *Handler) AddDriverWithVehicle(req *http.Request) handler.Response {

	var request DriverWithVehicleReq
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	err = h.ctrl.AddDriverWithVehicle(request)
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
			"data": "driver with vehicle created successfully",
		},
	}
}

func (h *Handler) GetDriverHistory(req *http.Request) handler.Response {
	driverId, ok := mux.Vars(req)["driverId"]
	if !ok || driverId == "" {
		return handler.BadRequest("invalid driverId")
	}

	id, _ := strconv.Atoi(driverId)
	resp, err := h.ctrl.GetDriverHistory(id)
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
