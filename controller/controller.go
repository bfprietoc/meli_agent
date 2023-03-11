package controller

import (
	"meli/config"
	"meli/model"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	db.Ping()
}

func GetAllData() ([]model.Data, error) {
	db := config.Connect()
	defer db.Close()

	var (
		data    model.Data
		arrData []model.Data
	)

	sqlStatement := `SELECT id, cpu_info, process_info, users_info, os_info FROM server_data`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.CpuInfo, &data.ProcessInfo, &data.UsersInfo, &data.OsInfo)
		if err != nil {
			return nil, err
		}
		arrData = append(arrData, data)
	}

	return arrData, nil

}

func InsertData(data model.Data) error {
	db := config.Connect()
	defer db.Close()

	sqlStatement := `INSERT INTO server_data(cpu_info, process_info, users_info, os_info) VALUES(?, ?, ?, ?)`

	_, err := db.Exec(sqlStatement, data.CpuInfo, data.ProcessInfo, data.UsersInfo, data.OsInfo)

	if err != nil {
		return err
	}

	return nil

}
