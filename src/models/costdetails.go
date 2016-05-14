package models

type ServerCost struct{
	RegionCost map[string]map[string]float64
}

type ServerResponse struct{
	Status bool `json:"status"`
	ErrorCode int `json:"code"`
	Message string `json:"message"`
	UserCost []CalculatedCost `json:"cost_details"`
}

type CalculatedCost struct {
	Region string `json:"region"`
	TotalCost float64 `json:"total_cost"`
	NumberOfServers int `json:"number_of_cpus"`
	ServerDetails map[string]int `json:"serverdetails"`
}