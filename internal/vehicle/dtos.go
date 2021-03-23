package vehicle

type CreateVehicleRequest struct {
	Model              string   `json:"model"`
	RegistrationNo     string   `json:"registration_no"`
	PermittedRideTypes []string `json:"permitted_ride_types"`
}
