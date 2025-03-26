package infrastructure

import (
	"log"

	"api2-multi.com/a/src/Notifications/domain"
	"api2-multi.com/a/src/core"
)

type ValidationMySQL struct {
	conn core.ConectionMySQL
}

func NewValidationMySQL()*ValidationMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&ValidationMySQL{conn: *conn}
}

func(r *ValidationMySQL)GetParametersOfCultivation(id_parcel int)([]domain.Parameters, error){
	query := "SELECT cultivation_parameters.humidity_min, cultivation_parameters.humidity_max, cultivation_parameters.temp_min, cultivation_parameters.temp_max, cultivation_parameters.min_air_con, cultivation_parameters.max_air_con FROM cultivation_parameters INNER JOIN crop ON cultivation_parameters.id_cultivation_parameter = crop.id_cultivation_parameter INNER JOIN parcel ON parcel.id_crop = crop.id_crop WHERE parcel.id_parcel = ?"
	rows, err := r.conn.FetchRows(query, id_parcel)
	var parameters []domain.Parameters
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var hum_min float32
		var hum_max float32
		var temp_min float32
		var temp_max float32
		var air_min float32
		var air_max float32

		if err := rows.Scan(&hum_min, &hum_max, &temp_min, &temp_max, &air_min, &air_max); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		parameter := &domain.Parameters{
			Min_hum: hum_min,
			Max_hum: hum_max,
			Min_temp: temp_min,
			Max_temp: temp_max,
			Min_air: air_min,
			Max_air: air_max,
		}
		parameters = append(parameters, *parameter)
	}
	return parameters, err
}