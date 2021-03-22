package address

type Location struct {
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Name       string  `json:"name"`
	StreetName string  `json:"street_name"`
	Landmark   string  `json:"landmark"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
}
