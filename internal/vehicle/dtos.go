package vehicle

import "github.com/pkg/errors"

var permittedRideTypes = map[string]bool{
	"MICRO":    true,
	"MINI":     true,
	"SEDAN":    true,
	"RACE_CAR": false,
}

type CreateVehicleRequest struct {
	Model              string   `json:"model"`
	RegistrationNo     string   `json:"registration_no"`
	PermittedRideTypes []string `json:"permitted_ride_types"`
}

func (c *CreateVehicleRequest) Validate() error {
	if c.Model == "" {
		return errors.Errorf("invalid ride model")
	}
	if c.RegistrationNo == "" {
		return errors.Errorf("invalid ride registration_no")
	}
	if len(c.PermittedRideTypes) == 0 {
		return errors.Errorf("permitted_ride_types cannot be empty")
	}
	for _, v := range c.PermittedRideTypes {
		if val, ok := permittedRideTypes[v]; !ok || val == false {
			return errors.Errorf("invalid ride type, %v", v)
		}
	}

	return nil
}
