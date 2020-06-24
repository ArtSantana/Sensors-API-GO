package repository

import (
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
