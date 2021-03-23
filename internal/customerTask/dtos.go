package customerTask

import (
	"github.com/atulsinha007/uber/internal/address"
)

type CustomerTask struct {
	CustomerTaskId int  `json:"customer_task_id"`
	CustomerId     int  `json:"customer_id"`
	Status         string  `json:"status"`
	PayableAmount  float64 `json:"payable_amount"`
	RideType       string  `json:"ride_type"`
}

type CustomerHistoryResponse struct {
	RideId        int             `json:"ride_id"`
	RideStops     []address.Location `json:"ride_stops"`
	Status        string             `json:"status"`
	PayableAmount float64            `json:"payable_amount"`
	PaymentStatus string             `json:"payment_status"`
	DriverInfo    DriverInfo         `json:"driver_info"`
	RatingGiven   string             `json:"rating_given"`
	DateOfJourney string             `json:"date_of_journey"`
}

type DriverInfo struct {
	Name    string `json:"name"`
	PhoneNo string `json:"phone_no"`
}

type CreateRideRequest struct {
	CustomerId        int           `json:"customer_id"`
	PayableAmount     float64          `json:"payable_amount"`
	PickupLocation    address.Location `json:"pickup_location"`
	DropLocation      address.Location `json:"drop_location"`
	PreferredRideType string           `json:"preferred_ride_type"`
}

type CreateRideResponseOnDriverAcceptance struct {
	PickupLocation address.Location `json:"pickup_location"`
	ETA            float64          `json:"eta"`
}

type UpdateRideReq struct {
	CustomerTaskId int
	Stops          []address.Location
}
