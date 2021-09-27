package models

type Ship struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Weight float32 `json:"weight"`
}
