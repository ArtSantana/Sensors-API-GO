package controller

import (
	"encoding/json"
	"net/http"
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
