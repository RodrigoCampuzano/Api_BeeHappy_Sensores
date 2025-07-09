package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
    var errStr string

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    dbHost   := os.Getenv("DB_HOST")
    dbUser   := os.Getenv("DB_USER")
    dbPass   := os.Getenv("DB_PASS")
    dbSchema := os.Getenv("DB_NAME")

    // Añadimos parseTime=true para que DATETIME/TIMESTAMP se mapeen a time.Time,
    // loc=Local para usar la zona horaria local.
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local",
        dbUser, dbPass, dbHost, dbSchema,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        errStr = fmt.Sprintf("error al abrir la base de datos: %v", err)
    }

    // Configuración del pool de conexiones
    db.SetMaxOpenConns(10)

    // Probar la conexión
    if pingErr := db.Ping(); pingErr != nil {
        db.Close()
        errStr = fmt.Sprintf("error al verificar la conexión a la base de datos: %v", pingErr)
    }

    return &Conn_MySQL{DB: db, Err: errStr}
}


func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) *sql.Rows {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta SELECT: %v", err)
	}

	return rows
}
