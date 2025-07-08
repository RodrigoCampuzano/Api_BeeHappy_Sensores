package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"database/sql"
	"fmt"
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

	var colmenaRaspberry entities.ColmenaRaspberryPi
	query := `SELECT id, id_colmena, id_raspberry, fecha_asignacion, estado 
		FROM colmena_raspberry WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&colmenaRaspberry.ID,
		&colmenaRaspberry.ID_Colmena,
		&colmenaRaspberry.ID_Raspberry,
		&colmenaRaspberry.Fecha_Asignacion,
		&colmenaRaspberry.Estado)

	if err != nil {
		if err == sql.ErrNoRows {
			return colmenaRaspberry, fmt.Errorf("relación colmena-raspberry no encontrada")
		}
		log.Printf("Error obteniendo colmena-raspberry: %v", err)
		return colmenaRaspberry, err
	}

	return colmenaRaspberry, nil
}

func (mysql *ColmenaRaspberryMySQL) UpdateColmenaRaspberry(cr entities.ColmenaRaspberryPi) error {

	// Primero verificamos si existe
	checkQuery := `SELECT id FROM colmena_raspberry WHERE id = ?`
	var existingID int
	err := mysql.conn.DB.QueryRow(checkQuery, cr.ID).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("relación colmena-raspberry no encontrada")
		}
		return err
	}

	// Si existe, actualizamos
	query := `
        UPDATE colmena_raspberry 
        SET id_colmena = ?,
            id_raspberry = ?,
            estado = ?
        WHERE id = ?
    `

	result, err := mysql.conn.DB.Exec(query,
		cr.ID_Colmena,
		cr.ID_Raspberry,
		cr.Estado,
		cr.ID)

	if err != nil {
		log.Printf("Error actualizando colmena-raspberry: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se actualizó ninguna relación colmena-raspberry")
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
