package user

import (
	"database/sql"
	"errors"
	"fmt"
)

// PostgresUserRepository es una implementación del UserRepository que utiliza PostgreSQL como almacenamiento.
type PostgresUserRepository struct {
	DB *sql.DB
}

// CreateUser crea un nuevo usuario en la base de datos PostgreSQL.
func (repo *PostgresUserRepository) CreateUser(user *User) error {
	// Definimos la consulta SQL para insertar un nuevo usuario en la tabla 'users'.
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"

	// Ejecutamos la consulta SQL con los datos del nuevo usuario y capturamos el resultado.
	_, err := repo.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		// En caso de error, se imprime un mensaje para identificar la causa del problema.
		fmt.Println("Error al insertar usuario:", err)
		return err
	}
	return nil
}

// FindByUsername busca un usuario en la base de datos PostgreSQL por su nombre de usuario.
func (repo *PostgresUserRepository) FindByUsername(username string) (*User, error) {
	// Definimos la consulta SQL para buscar un usuario por su nombre de usuario.
	query := "SELECT id, username, password FROM users WHERE username = $1"

	// Ejecutamos la consulta SQL y capturamos el resultado en una fila.
	row := repo.DB.QueryRow(query, username)

	// Creamos una variable para almacenar los datos del usuario.
	var user User

	// Escaneamos los datos de la fila en la variable 'user'.
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		// Si no se encuentra un usuario con el nombre de usuario proporcionado,
		// devolvemos un error con un mensaje indicando que el usuario no fue encontrado.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Si se encuentra un usuario con el nombre de usuario proporcionado,
	// devolvemos el usuario encontrado.
	return &user, nil
}
