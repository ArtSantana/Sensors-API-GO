package models

type Sensor struct {
	ID             int `json:"id"`
	Temperature    int `json:"temperature"`
	Active         int `json:"active"`
	MaxTemperature int `json:"max_temperature"`
}

type SensorNoID struct {
	Temperature    int `json:"temperature"`
	Active         int `json:"active"`
	MaxTemperature int `json:"max_temperature"`
}
