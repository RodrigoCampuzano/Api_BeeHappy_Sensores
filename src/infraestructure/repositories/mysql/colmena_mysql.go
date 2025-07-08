package mysql

import (
	core "apisensores/src/core/db"
	"apisensores/src/domain/entities"
	"database/sql"
	"fmt"
	"log"
)

type ColmenaMySQL struct {
	conn *core.Conn_MySQL
}

func NewColmenaMySQL() *ColmenaMySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error connecting to MySQL database: %v", conn.Err)
	}
	return &ColmenaMySQL{conn: conn}
}

func (mysql *ColmenaMySQL) CreateColmena(colmena entities.Colmenas) error {
	query := `INSERT INTO colmenas (identificador, nombre, area_ubicacion, tipo_colmena, estado) 
		VALUES (?, ?, ?, ?, ?)`

	if colmena.Estado == "" {
		colmena.Estado = "activo"
	}

	_, err := mysql.conn.DB.Exec(query,
		colmena.Identificador,
		colmena.Nombre,
		colmena.Area_Ubicacion,
		colmena.Tipo_Colmena,
		colmena.Estado)

	if err != nil {
		log.Printf("Error creating colmena: %v", err)
		return err
	}

	return nil
}

func (mysql *ColmenaMySQL) GetColmena(id int) (entities.Colmenas, error) {
	var colmena entities.Colmenas

	query := `SELECT id, identificador, nombre, area_ubicacion, tipo_colmena, estado, 
		fecha_registro, fecha_actualizacion FROM colmenas WHERE id = ?`

	err := mysql.conn.DB.QueryRow(query, id).Scan(
		&colmena.ID,
		&colmena.Identificador,
		&colmena.Nombre,
		&colmena.Area_Ubicacion,
		&colmena.Tipo_Colmena,
		&colmena.Estado,
		&colmena.Fecha_Registro,
		&colmena.Fecha_Actualizacion)

	if err != nil {
		log.Printf("Error getting colmena: %v", err)
		return colmena, err
	}

	return colmena, nil
}

func (mysql *ColmenaMySQL) UpdateColmena(colmena entities.Colmenas) error {
	// Primero verificamos si la colmena existe
	checkQuery := `SELECT id FROM colmenas WHERE id = ?`
	var existingID int
	err := mysql.conn.DB.QueryRow(checkQuery, colmena.ID).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("colmena no encontrada")
		}
		return err
	}

	// Si existe, procedemos a actualizar
	query := `
        UPDATE colmenas 
        SET identificador = ?, 
            nombre = ?, 
            area_ubicacion = ?, 
            tipo_colmena = ?, 
            estado = ?,
            fecha_actualizacion = CURRENT_TIMESTAMP
        WHERE id = ?
    `

	result, err := mysql.conn.DB.Exec(query,
		colmena.Identificador,
		colmena.Nombre,
		colmena.Area_Ubicacion,
		colmena.Tipo_Colmena,
		colmena.Estado,
		colmena.ID)

	if err != nil {
		log.Printf("Error actualizando colmena: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se actualizó ninguna colmena")
	}

	return nil
}

func (mysql *ColmenaMySQL) DeleteColmena(id int) error {
	// Primero verificamos si la colmena existe
	query := `SELECT id FROM colmenas WHERE id = ?`
	var existingID int
	err := mysql.conn.DB.QueryRow(query, id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("colmena no encontrada")
		}
		return err
	}

	// Si existe, procedemos a eliminar
	deleteQuery := `DELETE FROM colmenas WHERE id = ?`
	_, err = mysql.conn.DB.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Error al eliminar colmena: %v", err)
		return err
	}

	return nil
}

func (mysql *ColmenaMySQL) GetAllColmenas() ([]entities.Colmenas, error) {
	var colmenas []entities.Colmenas

	query := `SELECT id, identificador, nombre, area_ubicacion, tipo_colmena, estado, 
		fecha_registro, fecha_actualizacion FROM colmenas`

	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Printf("Error getting all colmenas: %v", err)
		return colmenas, err
	}
	defer rows.Close()

	for rows.Next() {
		var colmena entities.Colmenas
		err := rows.Scan(
			&colmena.ID,
			&colmena.Identificador,
			&colmena.Nombre,
			&colmena.Area_Ubicacion,
			&colmena.Tipo_Colmena,
			&colmena.Estado,
			&colmena.Fecha_Registro,
			&colmena.Fecha_Actualizacion)
		if err != nil {
			log.Printf("Error scanning colmena: %v", err)
			return colmenas, err
		}
		colmenas = append(colmenas, colmena)
	}

	return colmenas, nil
}

func (mysql *ColmenaMySQL) UpdateEstadoColmena(id int, estado string) error {
	// Primero verificamos si la colmena existe
	checkQuery := `SELECT id FROM colmenas WHERE id = ?`
	var existingID int
	err := mysql.conn.DB.QueryRow(checkQuery, id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("colmena no encontrada")
		}
		return err
	}

	// Validamos el estado
	estadosValidos := []string{"activo", "inactivo"}
	estadoValido := false
	for _, e := range estadosValidos {
		if e == estado {
			estadoValido = true
			break
		}
	}
	if !estadoValido {
		return fmt.Errorf("estado no válido. Debe ser 'activo' o 'inactivo'")
	}

	// Si existe y el estado es válido, procedemos a actualizar
	query := `
        UPDATE colmenas 
        SET estado = ?,
            fecha_actualizacion = CURRENT_TIMESTAMP
        WHERE id = ?
    `

	result, err := mysql.conn.DB.Exec(query, estado, id)
	if err != nil {
		log.Printf("Error actualizando estado de colmena: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se actualizó ninguna colmena")
	}

	return nil
}
