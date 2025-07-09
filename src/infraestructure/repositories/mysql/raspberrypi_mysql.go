package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"database/sql"
	"fmt"
	"log"
)

type RaspberryPiMySQL struct {
	conn *core.Conn_MySQL
}

func NewRaspberryPiMySQL() *RaspberryPiMySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error connecting to MySQL database: %v", conn.Err)
	}
	return &RaspberryPiMySQL{conn: conn}
}

func (mysql *RaspberryPiMySQL) CreateRaspberryPi(rpi entities.Raspberrypi) error {
	query := `INSERT INTO raspberry_pi 
		(mac, ip_address, nombre, modelo, estado) 
		VALUES (?, ?, ?, ?, ?)`

	result, err := mysql.conn.DB.Exec(query,
		rpi.Mac,
		rpi.IP_Address,
		rpi.Nombre,
		rpi.Modelo,
		rpi.Estado)

	if err != nil {
		log.Printf("Error al crear raspberry pi: %v", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	log.Printf("Raspberry Pi creado con ID: %d", id)
	return nil
}

func (mysql *RaspberryPiMySQL) GetRaspberryPi(id int) (entities.Raspberrypi, error) {
	var raspberry entities.Raspberrypi
	var fechaUltimaConexion sql.NullString

	query := `SELECT id, mac, ip_address, nombre, modelo, estado, 
		fecha_registro, fecha_ultima_conexion 
		FROM raspberry_pi WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&raspberry.ID,
		&raspberry.Mac,
		&raspberry.IP_Address,
		&raspberry.Nombre,
		&raspberry.Modelo,
		&raspberry.Estado,
		&raspberry.Fecha_Registro,
		&fechaUltimaConexion)

	if err != nil {
		if err == sql.ErrNoRows {
			return raspberry, fmt.Errorf("raspberry pi no encontrado")
		}
		return raspberry, err
	}

	// Manejar el campo NULL
	if fechaUltimaConexion.Valid {
		raspberry.Fecha_Ultima_Conexion = fechaUltimaConexion.String
	}

	return raspberry, nil
}

func (mysql *RaspberryPiMySQL) UpdateRaspberryPi(id int, rpi entities.Raspberrypi) error {
	// Verificar si existe
	existQuery := `SELECT id FROM raspberry_pi WHERE id = ?`
	var existingID int
	err := mysql.conn.DB.QueryRow(existQuery, id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("raspberry pi con ID %d no encontrado", id)
		}
		return err
	}

	// Realizar la actualización
	query := `
		UPDATE raspberry_pi 
		SET mac = ?,
			ip_address = ?,
			nombre = ?,
			modelo = ?,
			estado = ?,
			fecha_ultima_conexion = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	result, err := mysql.conn.DB.Exec(query,
		rpi.Mac,
		rpi.IP_Address,
		rpi.Nombre,
		rpi.Modelo,
		rpi.Estado,
		id)

	if err != nil {
		log.Printf("Error actualizando raspberry pi: %v", err)
		return fmt.Errorf("error al actualizar raspberry pi: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se realizó ninguna actualización")
	}

	log.Printf("Raspberry Pi ID %d actualizado exitosamente", id)
	return nil
}

func (mysql *RaspberryPiMySQL) DeleteRaspberryPi(id int) error {
	query := `DELETE FROM raspberry_pi WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting raspberry_pi: %v", err)
		return err
	}

	return nil
}

func (mysql *RaspberryPiMySQL) GetAllRaspberryPi() ([]entities.Raspberrypi, error) {
	var raspberries []entities.Raspberrypi

	query := `SELECT id, mac, ip_address, nombre, modelo, estado, 
		fecha_registro, fecha_ultima_conexion 
		FROM raspberry_pi`

	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var raspberry entities.Raspberrypi
		var fechaUltimaConexion sql.NullString

		err := rows.Scan(
			&raspberry.ID,
			&raspberry.Mac,
			&raspberry.IP_Address,
			&raspberry.Nombre,
			&raspberry.Modelo,
			&raspberry.Estado,
			&raspberry.Fecha_Registro,
			&fechaUltimaConexion)

		if err != nil {
			return nil, err
		}

		if fechaUltimaConexion.Valid {
			raspberry.Fecha_Ultima_Conexion = fechaUltimaConexion.String
		}

		raspberries = append(raspberries, raspberry)
	}

	return raspberries, nil
}	

func (mysql *RaspberryPiMySQL) UpdateEstadoRaspberryPi(id int, estado string) error {
	query := `UPDATE raspberry_pi SET estado = ? WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, estado, id)
	if err != nil {
		log.Printf("Error updating estado raspberry_pi: %v", err)
		return err
	}
	
	return nil
}		

func (mysql *RaspberryPiMySQL) GetByMAC(mac string) (entities.Raspberrypi, error) {
	var raspberry entities.Raspberrypi
	var fechaUltimaConexion sql.NullString

	query := `SELECT id, mac, ip_address, nombre, modelo, estado, 
        fecha_registro, fecha_ultima_conexion 
        FROM raspberry_pi WHERE mac = ?`

	err := mysql.conn.DB.QueryRow(query, mac).Scan(
		&raspberry.ID,
		&raspberry.Mac,
		&raspberry.IP_Address,
		&raspberry.Nombre,
		&raspberry.Modelo,
		&raspberry.Estado,
		&raspberry.Fecha_Registro,
		&fechaUltimaConexion)

	if err != nil {
		if err == sql.ErrNoRows {
			return raspberry, fmt.Errorf("raspberry pi con MAC %s no encontrado", mac)
		}
		log.Printf("Error al obtener raspberry pi por MAC: %v", err)
		return raspberry, err
	}

	// Solo asignar el valor si no es NULL
	if fechaUltimaConexion.Valid {
		raspberry.Fecha_Ultima_Conexion = fechaUltimaConexion.String
	} else {
		raspberry.Fecha_Ultima_Conexion = "" // o cualquier valor por defecto que prefieras
	}

	return raspberry, nil
}
	