package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"log"
)

type TipoSensoresMySQL struct {
	conn *core.Conn_MySQL
}

func NewTipoSensoresMySQL() *TipoSensoresMySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error connecting to MySQL database: %v", conn.Err)
	}
	return &TipoSensoresMySQL{conn: conn}
}

func (mysql *TipoSensoresMySQL) CreateTipoSensor(tipoSensor entities.Tipo_Sesnores) error {
	query := `INSERT INTO tipos_sensores (nombre, descripcion, unidad_medida, tipo_dato) 
		VALUES (?, ?, ?, ?)`

	if tipoSensor.Tipo_Dato == "" {
		tipoSensor.Tipo_Dato = "decimal"
	}

	_, err := mysql.conn.DB.Exec(query,
		tipoSensor.Nombre,
		tipoSensor.Descripcion,
		tipoSensor.Unidad_Medida,
		tipoSensor.Tipo_Dato)

	if err != nil {
		log.Printf("Error creating tipo_sensor: %v", err)
		return err
	}

	return nil
}

func (mysql *TipoSensoresMySQL) GetTipoSensor(id int) (entities.Tipo_Sesnores, error) {
	var tipoSensor entities.Tipo_Sesnores

	query := `SELECT id, nombre, descripcion, unidad_medida, tipo_dato 
		FROM tipos_sensores WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&tipoSensor.ID,
		&tipoSensor.Nombre,
		&tipoSensor.Descripcion,
		&tipoSensor.Unidad_Medida,
		&tipoSensor.Tipo_Dato)

	if err != nil {
		log.Printf("Error getting tipo_sensor: %v", err)
		return tipoSensor, err
	}

	return tipoSensor, nil
}

func (mysql *TipoSensoresMySQL) UpdateTipoSensor(tipoSensor entities.Tipo_Sesnores) error {
	query := `UPDATE tipos_sensores SET nombre = ?, descripcion = ?, unidad_medida = ?, 
		tipo_dato = ? WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query,
		tipoSensor.Nombre,
		tipoSensor.Descripcion,
		tipoSensor.Unidad_Medida,
		tipoSensor.Tipo_Dato,
		tipoSensor.ID)

	if err != nil {
		log.Printf("Error updating tipo_sensor: %v", err)
		return err
	}

	return nil
}

func (mysql *TipoSensoresMySQL) DeleteTipoSensor(id int) error {
	query := `DELETE FROM tipos_sensores WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting tipo_sensor: %v", err)
		return err
	}

	return nil
}

func (mysql *TipoSensoresMySQL) GetAllTipoSensores() ([]entities.Tipo_Sesnores, error) {
	var tiposSensores []entities.Tipo_Sesnores

	query := `SELECT id, nombre, descripcion, unidad_medida, tipo_dato FROM tipos_sensores`

	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Printf("Error getting all tipos_sensores: %v", err)
		return tiposSensores, err
	}
	defer rows.Close()

	for rows.Next() {
		var tipoSensor entities.Tipo_Sesnores
		err := rows.Scan(
			&tipoSensor.ID,
			&tipoSensor.Nombre,
			&tipoSensor.Descripcion,
			&tipoSensor.Unidad_Medida,	
			&tipoSensor.Tipo_Dato)
		if err != nil {
			log.Printf("Error scanning tipos_sensores: %v", err)
			return tiposSensores, err
		}
		tiposSensores = append(tiposSensores, tipoSensor)
	}

	return tiposSensores, nil
}

func (mysql *TipoSensoresMySQL) GetByNombre(nombre string) (entities.Tipo_Sesnores, error) {
	var tipoSensor entities.Tipo_Sesnores

	query := `SELECT id, nombre, descripcion, unidad_medida, tipo_dato FROM tipos_sensores WHERE nombre = ?`

	err := mysql.conn.DB.QueryRow(query, nombre).Scan(
		&tipoSensor.ID,
		&tipoSensor.Nombre,
		&tipoSensor.Descripcion,
		&tipoSensor.Unidad_Medida,
		&tipoSensor.Tipo_Dato)

	if err != nil {
		log.Printf("Error getting tipo_sensor by nombre: %v", err)
		return tipoSensor, err
	}

	return tipoSensor, nil
}	