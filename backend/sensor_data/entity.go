package sensordata

import "time"

type SensorData struct {
	ID                int64     `json:"id"`
	FirstTemperature  float64   `json:"first_temperature"`
	SecondTemperature float64   `json:"second_temperature"`
	Humidity          float64   `json:"humidity"`
	CreatedAt         time.Time `json:"created_at"`
}

type SensorDataInput struct {
	FirstTemperature  float64 `json:"first_temperature"`
	SecondTemperature float64 `json:"second_temperature"`
	Humidity          float64 `json:"humidity"`
}
