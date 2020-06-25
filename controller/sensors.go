package controller

import (
	"encoding/json"
	"net/http"
	"sensores/models"
	"sensores/repository"
)

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
