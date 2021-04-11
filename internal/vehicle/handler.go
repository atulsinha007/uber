package vehicle

import (
	"encoding/json"
	handler "github.com/atulsinha007/uber/pkg/http-wrapper"
	"net/http"
)

type Handler struct {
	ctrl Ctrl
}

func NewHandler(ctrl Ctrl) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) CreateVehicle(req *http.Request) handler.Response {
	var payload CreateVehicleRequest
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return handler.BadRequest("invalid payload")
	}

	err = payload.Validate()
	if err != nil {
		return handler.BadRequest(err.Error())
	}

	err = h.ctrl.CreateVehicle(payload)
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
			"data": "vehicle creation successful",
		},
	}
}
