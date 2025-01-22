package sensordata

import "time"

type SensorData struct {
	ID                int64     `json:"id"`
	RoomTemperature   float64   `json:"room_temperature"`
	ObjectTemperature float64   `json:"object_temperature"`
	Humidity          float64   `json:"humidity"`
	CreatedAt         time.Time `json:"created_at"`
}

type SensorDataInput struct {
	RoomTemperature   float64 `json:"room_temperature"`
	ObjectTemperature float64 `json:"object_temperature"`
	Humidity          float64 `json:"humidity"`
}
