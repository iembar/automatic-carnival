package models



type FatalError struct{
	Status bool `json:"status"`
	ErrorCode int `json:"code"`
	Message string `json:"message"`
}
