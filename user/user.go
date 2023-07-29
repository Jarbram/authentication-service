package user

// Definición de la estructura User que representa un usuario.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Definición de la interfaz UserRepository que especifica los métodos que deben ser implementados
// por cualquier repositorio que desee trabajar con usuarios.
type UserRepository interface {
	// CreateUser crea un nuevo usuario en el repositorio.
	CreateUser(user *User) error

	// FindByUsername busca un usuario por su nombre de usuario y devuelve el usuario encontrado.
	// También puede devolver un error si el usuario no existe en el repositorio.
	FindByUsername(username string) (*User, error)
}
