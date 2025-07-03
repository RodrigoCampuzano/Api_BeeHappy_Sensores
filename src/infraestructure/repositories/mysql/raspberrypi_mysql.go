package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"log"
	"time"
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
	query := `INSERT INTO raspberry_pi (mac, nombre, modelo, ip_address, estado) 
		VALUES (?, ?, ?, ?, ?)`

	if rpi.Estado == "" {
		rpi.Estado = "activo"
	}

	_, err := mysql.conn.DB.Exec(query,
		rpi.Mac,
		rpi.Nombre,
		rpi.Modelo,
		rpi.IP_Address,
		rpi.Estado)

	if err != nil {
		log.Printf("Error creating raspberry_pi: %v", err)
		return err
	}

	return nil
}

func (mysql *RaspberryPiMySQL) GetRaspberryPi(id int) (entities.Raspberrypi, error) {
	var rpi entities.Raspberrypi

	query := `SELECT id, mac, nombre, modelo, ip_address, estado, fecha_registro, 
		fecha_ultima_conexion FROM raspberry_pi WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&rpi.ID,
		&rpi.Mac,
		&rpi.Nombre,
		&rpi.Modelo,
		&rpi.IP_Address,
		&rpi.Estado,
		&rpi.Fecha_Registro,
		&rpi.Fecha_Ultima_Conexion)

	if err != nil {
		log.Printf("Error getting raspberry_pi: %v", err)
		return rpi, err
	}

	return rpi, nil
}

func (mysql *RaspberryPiMySQL) UpdateRaspberryPi(rpi entities.Raspberrypi) error {
	query := `UPDATE raspberry_pi SET mac = ?, nombre = ?, modelo = ?, ip_address = ?, 
		estado = ?, fecha_ultima_conexion = ? WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query,
		rpi.Mac,
		rpi.Nombre,
		rpi.Modelo,
		rpi.IP_Address,
		rpi.Estado,
		time.Now(),
		rpi.ID)

	if err != nil {
		log.Printf("Error updating raspberry_pi: %v", err)
		return err
	}

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
	var rpis []entities.Raspberrypi

	query := `SELECT id, mac, nombre, modelo, ip_address, estado, fecha_registro, 
		fecha_ultima_conexion FROM raspberry_pi`

	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Printf("Error getting all raspberry_pi: %v", err)
		return rpis, err
	}
	defer rows.Close()

	for rows.Next() {
		var rpi entities.Raspberrypi
		err := rows.Scan(
			&rpi.ID,
			&rpi.Mac,
			&rpi.Nombre,
			&rpi.Modelo,
			&rpi.IP_Address,
			&rpi.Estado,
			&rpi.Fecha_Registro,
			&rpi.Fecha_Ultima_Conexion)		
		if err != nil {
			log.Printf("Error scanning raspberry_pi: %v", err)
			return rpis, err
		}
		rpis = append(rpis, rpi)
	}

	return rpis, nil
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

func (mysql *RaspberryPiMySQL) GetRaspberryPiByMac(mac string) ([]*entities.Raspberrypi, error) {
	var rpis []*entities.Raspberrypi

	query := `SELECT id, mac, nombre, modelo, ip_address, estado, fecha_registro, 
		fecha_ultima_conexion FROM raspberry_pi WHERE mac = ?`

	rows, err := mysql.conn.DB.Query(query, mac)
	if err != nil {
		log.Printf("Error getting raspberry_pi by mac: %v", err)
		return rpis, err
	}
	defer rows.Close()

	for rows.Next() {
		var rpi entities.Raspberrypi
		err := rows.Scan(
		&rpi.ID,
		&rpi.Mac,	
		&rpi.Nombre,
		&rpi.Modelo,
		&rpi.IP_Address,
		&rpi.Estado,
		&rpi.Fecha_Registro,
		&rpi.Fecha_Ultima_Conexion)

	if err != nil {
		log.Printf("Error getting raspberry_pi by mac: %v", err)		
		return rpis, err
	}
	rpis = append(rpis, &rpi)	
	}

	return rpis, nil
}

	