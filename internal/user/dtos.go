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
	DriverId  string
	VehicleId string
}

type DriverProfileResponse struct {
	Name          string
	PhoneNo       string
	TotalRides    int
	AverageRating string
}

type DriverHistoryRequest struct {
	DriverId string
}

type DriverHistoryResponse struct {
	RideId          string
	Status          string
	DistanceCovered float64
	Rating          int
	PayoutAmount    float64
}

type UpdateRideRequest struct {
	DriverTaskId string
	Status       string
}

type UpdateCurrentLocationRequest struct {
	UserId      string
	CurLocation LatLng
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
