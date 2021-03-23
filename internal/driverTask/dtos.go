package driverTask

type DriverHistoryResponse struct {
	RideId          int     `json:"ride_id"`
	Status          string  `json:"status"`
	DistanceCovered float64 `json:"distance_covered"`
	Rating          int     `json:"rating"`
	PayoutAmount    float64 `json:"payout_amount"`
}

type AcceptRideReq struct {
	DriverTaskId   int `json:"driver_task_id"`
	DriverId       int `json:"driver_id"`
	CustomerTaskId int `json:"customer_task_id"`
}

type UpdateRideReq struct {
	DriverTaskId   int    `json:"driver_task_id"`
	Status         string `json:"status"`
	CustomerTaskId int    `json:"customer_task_id"`
}

type DriverTask struct {
	DriverTaskId   int
	CustomerTaskId int
	DriverId       int
	Status         string
	PayableAmount  float64
	RideType       string
	Distance       float64
}