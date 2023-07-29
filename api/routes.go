package api

import (
	"authentication-service/authentication"
	"authentication-service/user"

	"github.com/gorilla/mux"
)

// SetupRoutes configura las rutas del API y asocia los handlers correspondientes.
// Recibe un enrutador (Router) de gorilla/mux, el repositorio de usuarios (userRepo) y el servicio de autenticación (authService).
func SetupRoutes(r *mux.Router, userRepo user.UserRepository, authService *authentication.AuthService) {
	// Configuración de la ruta para el registro de usuarios (/register) con el método POST.
	// Se asocia la función registerUserHandler al handler de esta ruta,
	// pasando el repositorio de usuarios y el servicio de autenticación como argumentos.
	r.HandleFunc("/register", registerUserHandler(userRepo, authService)).Methods("POST")

	// Configuración de la ruta para la autenticación de usuarios (/login) con el método POST.
	// Se asocia la función authenticateUserHandler al handler de esta ruta,
	// pasando el repositorio de usuarios y el servicio de autenticación como argumentos.
	r.HandleFunc("/login", authenticateUserHandler(userRepo, authService)).Methods("POST")
}
