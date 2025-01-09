package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jovi345/sensor-suhu-kelembapan/handler"
	sensordata "github.com/jovi345/sensor-suhu-kelembapan/sensor_data"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/temp_humid?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")

	sensorDataRepository := sensordata.NewRepository(db)
	sensorDataService := sensordata.NewService(sensorDataRepository)
	sensorDataHandler := handler.NewSensorDataHandler(sensorDataService)

	http.HandleFunc("/api/v1/data/add", sensorDataHandler.InsertData)
	http.HandleFunc("/api/v1/data/get", sensorDataHandler.GetAllData)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server: Failed to run!")
	}
}
