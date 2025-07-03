package mysql

import (
	"apisensores/src/domain/entities"
	core "apisensores/src/core/db"
	"log"
)

type ColmenaRaspberryMySQL struct {
	conn *core.Conn_MySQL
}

func NewColmenaRaspberryMySQL() *ColmenaRaspberryMySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error connecting to MySQL database: %v", conn.Err)
	}
	return &ColmenaRaspberryMySQL{conn: conn}
}

func (mysql *ColmenaRaspberryMySQL) CreateColmenaRaspberry(cr entities.ColmenaRaspberryPi) error {
	query := `INSERT INTO colmena_raspberry (id_colmena, id_raspberry, estado) 
		VALUES (?, ?, ?)`

	if cr.Estado == "" {
		cr.Estado = "activo"
	}

	_, err := mysql.conn.DB.Exec(query,
		cr.ID_Colmena,
		cr.ID_Raspberry,
		cr.Estado)

	if err != nil {
		log.Printf("Error creating colmena_raspberry: %v", err)
		return err
	}

	return nil
}

func (mysql *ColmenaRaspberryMySQL) GetColmenaRaspberry(id int) (entities.ColmenaRaspberryPi, error) {
	var cr entities.ColmenaRaspberryPi

	query := `SELECT id, id_colmena, id_raspberry, fecha_asignacion, estado 
		FROM colmena_raspberry WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&cr.ID,
		&cr.ID_Colmena,
		&cr.ID_Raspberry,
		&cr.Fecha_Asignacion,
		&cr.Estado)

	if err != nil {
		log.Printf("Error getting colmena_raspberry: %v", err)
		return cr, err
	}

	return cr, nil
}

func (mysql *ColmenaRaspberryMySQL) UpdateColmenaRaspberry(cr entities.ColmenaRaspberryPi) error {
	query := `UPDATE colmena_raspberry SET id_colmena = ?, id_raspberry = ?, estado = ? 
		WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query,
		cr.ID_Colmena,
		cr.ID_Raspberry,
		cr.Estado,
		cr.ID)

	if err != nil {
		log.Printf("Error updating colmena_raspberry: %v", err)
		return err
	}

	return nil
}

func (mysql *ColmenaRaspberryMySQL) DeleteColmenaRaspberry(id int) error {
	query := `DELETE FROM colmena_raspberry WHERE id = ?`

	_, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting colmena_raspberry: %v", err)
		return err
	}

	return nil
}
