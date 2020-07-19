package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sensores/global"
	"sensores/models"
	"sensores/repository"
)

// GetSensors return the sensors from database
func GetSensors(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	sensors := repository.GetSensors()

	result, err := json.Marshal(sensors)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error mashalling the sensors"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
	return
}

// UpdateSensorTemp update the temperature of sensors
func UpdateSensorTemp(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var updateSensor models.UpdateModel

	err := decoder.Decode(&updateSensor)

	if err != nil {
		panic(err.Error())
	}

	isError, updatedSensor := repository.UpdateTemp(updateSensor)

	if isError {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, errRes := json.Marshal(updatedSensor)

	if errRes != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error mashalling the sensors"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)

}

// CreateSensor controller
func CreateSensor(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var sensor models.Sensor

	err := decoder.Decode(&sensor)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	isError, createdSensor := repository.CreateSensor(sensor)

	if isError {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, errResult := json.Marshal(createdSensor)

	if errResult != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	uri := fmt.Sprintf("%s/sensors/%d", global.Env("HOST"), createdSensor.ID)

	res.Header().Add("Location", uri)
	res.Write(result)
}
