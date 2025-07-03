package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"log"
)

type SensoresMySQL struct {
	conn *core.Conn_MySQL
}

func NewSensoresMySQL() *SensoresMySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error connecting to MySQL database: %v", conn.Err)
	}
	return &SensoresMySQL{conn: conn}
}

func (mysql *SensoresMySQL) CreateSensor(sensor entities.Sensores) error {
	query := `INSERT INTO sensores (id_tipo_sensor, id_raspberry, pin_conexion, 
		direccion_i2c, estado) VALUES (?, ?, ?, ?, ?)`

	if sensor.Estado == "" {
		sensor.Estado = "activo"
	}

	_, err := mysql.conn.DB.Exec(query,
		sensor.ID_Tipo_Sensor,
		sensor.ID_Raspberry,
		sensor.Pin_conexion,
		sensor.Direccion_i2c,
		sensor.Estado)

	if err != nil {
		log.Printf("Error creating sensor: %v", err)
		return err
	}

	return nil
}

func (mysql *SensoresMySQL) GetSensor(id int) (entities.Sensores, error) {
	var sensor entities.Sensores

	query := `SELECT id, id_tipo_sensor, id_raspberry, pin_conexion, direccion_i2c, 
		estado, fecha_registro FROM sensores WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&sensor.ID,
		&sensor.ID_Tipo_Sensor,
		&sensor.ID_Raspberry,
		&sensor.Pin_conexion,
		&sensor.Direccion_i2c,
		&sensor.Estado,
		&sensor.Fecha_Registro)

	if err != nil {
		log.Printf("Error getting sensor: %v", err)
		return sensor, err
	}

	return sensor, nil
}

func (mysql *SensoresMySQL) UpdateSensor(sensor entities.Sensores) error {
	query := `UPDATE sensores SET id_tipo_sensor = ?, id_raspberry = ?, pin_conexion = ?, 
		direccion_i2c = ?, estado = ? WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query,
		sensor.ID_Tipo_Sensor,
		sensor.ID_Raspberry,
		sensor.Pin_conexion,
		sensor.Direccion_i2c,
		sensor.Estado,
		sensor.ID)

	if err != nil {
		log.Printf("Error updating sensor: %v", err)
		return err
	}

	return nil
}

func (mysql *SensoresMySQL) DeleteSensor(id int) error {
	query := `DELETE FROM sensores WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting sensor: %v", err)
		return err
	}

	return nil
}

func (mysql *SensoresMySQL) GetAllSensores() ([]entities.Sensores, error) {
	var sensores []entities.Sensores

	query := `SELECT id, id_tipo_sensor, id_raspberry, pin_conexion, direccion_i2c, 
		estado, fecha_registro FROM sensores`

	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Printf("Error getting all sensores: %v", err)
		return sensores, err
	}
	defer rows.Close()

	for rows.Next() {
		var sensor entities.Sensores
		err := rows.Scan(
			&sensor.ID,
			&sensor.ID_Tipo_Sensor,
			&sensor.ID_Raspberry,
			&sensor.Pin_conexion,
			&sensor.Direccion_i2c,
			&sensor.Estado,
			&sensor.Fecha_Registro)
		if err != nil {
			log.Printf("Error scanning sensores: %v", err)
			return sensores, err
		}
		sensores = append(sensores, sensor)
	}

	return sensores, nil
}

func (mysql *SensoresMySQL) GetByRaspberryID(id int) ([]entities.Sensores, error) {
	var sensores []entities.Sensores

	query := `SELECT id, id_tipo_sensor, id_raspberry, pin_conexion, direccion_i2c, 
		estado, fecha_registro FROM sensores WHERE id_raspberry = ?`

	rows, err := mysql.conn.DB.Query(query, id)
	
	if err != nil {
		log.Printf("Error getting sensores by raspberry: %v", err)
		return sensores, err
	}
	defer rows.Close()

	for rows.Next() {
		var sensor entities.Sensores
		err := rows.Scan(
			&sensor.ID,
			&sensor.ID_Tipo_Sensor,
			&sensor.ID_Raspberry,
			&sensor.Pin_conexion,
			&sensor.Direccion_i2c,
			&sensor.Estado,
			&sensor.Fecha_Registro)
		if err != nil {
			log.Printf("Error scanning sensores: %v", err)
			return sensores, err
		}
		sensores = append(sensores, sensor)
	}

	return sensores, nil
}

func (mysql *SensoresMySQL) UpdateEstadoSensor(id int, estado string) error {
	query := `UPDATE sensores SET estado = ? WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, estado, id)
	if err != nil {
		log.Printf("Error updating estado sensor: %v", err)
		return err
	}

	return nil
}	