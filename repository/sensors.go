package repository

import (
	"fmt"
	"log"
	"sensores/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetSensors() []models.Sensor {
	var (
		sensors []models.Sensor
	)

	queryString := "SELECT * FROM sensors"

	result, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var sensor models.Sensor

		err = result.Scan(
			&sensor.ID,
			&sensor.Temperature,
			&sensor.Active,
			&sensor.MaxTemperature,
		)

		if err != nil {
			panic(err.Error())
		}

		sensors = append(sensors, sensor)
	}

	return sensors
}

func UpdateTemp(updateSensor models.UpdateModel) (bool, models.Sensor) {
	var updatedSensor models.Sensor

	queryString := fmt.Sprintf("UPDATE sensors SET temperature = %d WHERE ID = %d", updateSensor.Temperature, updateSensor.ID)

	result, err := db.Query(queryString)

	if err != nil {
		log.Println(result)
		return true, updatedSensor
	}

	queryString = fmt.Sprintf("SELECT * FROM sensors WHERE ID = %d", updateSensor.ID)

	resultGet, errGet := db.Query(queryString)

	if errGet != nil {
		return true, updatedSensor
	}

	for resultGet.Next() {
		err = resultGet.Scan(
			&updatedSensor.ID,
			&updatedSensor.Temperature,
			&updatedSensor.Active,
			&updatedSensor.MaxTemperature,
		)

		if err != nil {
			return true, updatedSensor
		}
	}

	return false, updatedSensor
}
