package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Constantes para la configuración de la conexión con la base de datos.
const (
	dbHost     = "localhost"  // Dirección del servidor de la base de datos.
	dbPort     = "5432"       // Puerto de escucha de la base de datos.
	dbName     = "mydb"       // Nombre de la base de datos a la que se desea conectar.
	dbUser     = "myuser"     // Nombre de usuario para autenticarse en la base de datos.
	dbPassword = "mypassword" // Contraseña para autenticarse en la base de datos.
)

// NewDB crea y retorna una nueva instancia de conexión a la base de datos PostgreSQL.
func NewDB() (*sql.DB, error) {
	// Componer la cadena de conexión con los datos de configuración.
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Abrir la conexión con la base de datos.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Devuelve la instancia de conexión a la base de datos.
	return db, nil
}
