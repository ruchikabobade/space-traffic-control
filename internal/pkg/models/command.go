package models

type Command struct {
	Command string `json:"command"`
	DockingStation string `json:"dockingStation,omitempty"`
	Duration int `json:"duration,omitempty"`
}
