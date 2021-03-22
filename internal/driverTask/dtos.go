package driverTask

type DriverHistoryResponse struct {
	RideId          string
	Status          string
	DistanceCovered float64
	Rating          int
	PayoutAmount    float64
}

type AcceptRideReq struct {
	DriverTaskId string
	DriverId     string
}

type UpdateRideReq struct {
	DriverTaskId string
	Status     string
}
