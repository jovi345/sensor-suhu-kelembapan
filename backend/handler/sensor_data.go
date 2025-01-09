package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	sensordata "github.com/jovi345/sensor-suhu-kelembapan/sensor_data"
	"github.com/jovi345/sensor-suhu-kelembapan/utils"
)

type sensorDataHandler struct {
	sensorDataService sensordata.Service
}

func NewSensorDataHandler(sensorDataService sensordata.Service) *sensorDataHandler {
	return &sensorDataHandler{sensorDataService}
}

func (h *sensorDataHandler) InsertData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	var input sensordata.SensorDataInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	result, err := h.sensorDataService.InsertData(input)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, result)
}

func (h *sensorDataHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	result, err := h.sensorDataService.GetAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}
