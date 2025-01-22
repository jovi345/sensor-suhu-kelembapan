package sensordata

import (
	"time"
)

type Service interface {
	InsertData(input SensorDataInput) (SensorData, error)
	GetAll() ([]SensorData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InsertData(input SensorDataInput) (SensorData, error) {
	data := SensorData{
		RoomTemperature:   input.RoomTemperature,
		ObjectTemperature: input.ObjectTemperature,
		Humidity:          input.Humidity,
		CreatedAt:         time.Now(),
	}

	result, err := s.repository.Save(data)
	if err != nil {
		return SensorData{}, err
	}

	return result, nil
}

func (s *service) GetAll() ([]SensorData, error) {
	result, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}
