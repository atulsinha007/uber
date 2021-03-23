package user

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Type      string `json:"type"`
}

func CreateUserRequestToUser(req CreateUserRequest) User {
	return User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Type:      req.Type,
	}
}

type VehicleAssignmentRequest struct {
	DriverId  int
	VehicleId int
}

type DriverProfileResponse struct {
	Name          string
	PhoneNo       string
	TotalRides    int
	AverageRating string
}

type DriverHistoryRequest struct {
	driverId int
}

type DriverHistoryResponse struct {
	RideId          int
	Status          string
	DistanceCovered float64
	Rating          int
	PayoutAmount    float64
}

type UpdateRideRequest struct {
	DriverTaskId int
	Status       string
}

type UpdateCurrentLocationRequest struct {
	UserId      int
	CurLocation LatLng
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type DriverWithVehicleReq struct {
	CreateUserRequest
	Model              string   `json:"model"`
	RegistrationNo     string   `json:"registration_no"`
	PermittedRideTypes []string `json:"permitted_ride_types"`
}