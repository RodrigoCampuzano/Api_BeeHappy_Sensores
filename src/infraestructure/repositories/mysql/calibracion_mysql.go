package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"log"
	"time"
)

type CalibracionMySQL struct {
	conn *core.Conn_MySQL
}

func NewCalibracionMySQL() *CalibracionMySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error connecting to MySQL database: %v", conn.Err)
	}
	return &CalibracionMySQL{conn: conn}
}

func (mysql *CalibracionMySQL) CreateCalibracion(calibracion entities.Calibracion) error {
	query := `INSERT INTO calibracion (id_colmena, id_sensor, valor_minimo, valor_maximo, 
		factor_calibracion, offset_calibracion, fecha_calibracion) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	if calibracion.Fecha_Calibracion.IsZero() {
		calibracion.Fecha_Calibracion = time.Now()
	}

	_, err := mysql.conn.DB.Exec(query,
		calibracion.ID_Colmena,
		calibracion.ID_Sensor,
		calibracion.Valor_Minimo,
		calibracion.Valor_Maximo,
		calibracion.Factor_Calibracion,
		calibracion.Offset_Calibracion,
		calibracion.Fecha_Calibracion)

	if err != nil {
		log.Printf("Error creating calibracion: %v", err)
		return err
	}

	return nil
}

func (mysql *CalibracionMySQL) GetCalibracion(id int) (entities.Calibracion, error) {
	var calibracion entities.Calibracion

	query := `SELECT id, id_colmena, id_sensor, valor_minimo, valor_maximo, 
		factor_calibracion, offset_calibracion, fecha_calibracion 
		FROM calibracion WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&calibracion.ID,
		&calibracion.ID_Colmena,
		&calibracion.ID_Sensor,
		&calibracion.Valor_Minimo,
		&calibracion.Valor_Maximo,
		&calibracion.Factor_Calibracion,
		&calibracion.Offset_Calibracion,
		&calibracion.Fecha_Calibracion)

	if err != nil {
		log.Printf("Error getting calibracion: %v", err)
		return calibracion, err
	}

	return calibracion, nil
}

func (mysql *CalibracionMySQL) UpdateCalibracion(calibracion entities.Calibracion) error {
	query := `UPDATE calibracion SET id_colmena = ?, id_sensor = ?, valor_minimo = ?, 
		valor_maximo = ?, factor_calibracion = ?, offset_calibracion = ?, 
		fecha_calibracion = ? WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query,
		calibracion.ID_Colmena,
		calibracion.ID_Sensor,
		calibracion.Valor_Minimo,
		calibracion.Valor_Maximo,
		calibracion.Factor_Calibracion,
		calibracion.Offset_Calibracion,
		calibracion.Fecha_Calibracion,
		calibracion.ID)

	if err != nil {
		log.Printf("Error updating calibracion: %v", err)
		return err
	}

	return nil
}

func (mysql *CalibracionMySQL) DeleteCalibracion(id int) error {
	query := `DELETE FROM calibracion WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting calibracion: %v", err)
		return err
	}

	return nil
}
