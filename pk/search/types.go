package search

type LocationResult struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Responder struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
