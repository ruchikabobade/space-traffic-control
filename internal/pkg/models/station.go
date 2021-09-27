package models

type Station struct {
	ID string `json:"id,omitempty"`
	Capacity float32 `json:"capacity"`
	UsedCapacity float32 `json:"usedCapacity,omitempty"`
	Docks []Dock `json:"docks"`
}

type Dock struct {
	ID string `json:"id,omitempty"`
	NumDockingPorts int `json:"numDockingPorts"`
	Occupied int `json:"occupied,omitempty"`
	Weight float32 `json:"weight,omitempty"`
}