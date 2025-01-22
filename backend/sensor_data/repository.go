package sensordata

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(data SensorData) (SensorData, error)
	GetAll() ([]SensorData, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(data SensorData) (SensorData, error) {
	query := `
		INSERT INTO sensor_data 
		(room_temperature, object_temperature, humidity, created_at)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		context.Background(),
		query,
		data.RoomTemperature,
		data.ObjectTemperature,
		data.Humidity,
		data.CreatedAt,
	)
	if err != nil {
		return SensorData{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return SensorData{}, err
	}
	data.ID = id

	return data, nil
}

func (r *repository) GetAll() ([]SensorData, error) {
	var datas []SensorData

	query := `SELECT * FROM sensor_data`
	rows, err := r.db.Query(query)
	if err != nil {
		return datas, err
	}
	defer rows.Close()

	for rows.Next() {
		var data SensorData
		err = rows.Scan(
			&data.ID,
			&data.RoomTemperature,
			&data.ObjectTemperature,
			&data.Humidity,
			&data.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}
