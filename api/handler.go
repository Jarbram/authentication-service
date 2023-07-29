package api

import (
	"authentication-service/authentication"
	"authentication-service/user"
	"encoding/json"
	"net/http"
)

// registerUserHandler es el manejador HTTP para el endpoint de registro de usuarios.
// Recibe un repositorio de usuarios (userRepo) y un servicio de autenticación (authService).
func registerUserHandler(userRepo user.UserRepository, authService *authentication.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser user.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Intenta registrar al nuevo usuario utilizando el servicio de autenticación.
		err = authService.RegisterUser(newUser.Username, newUser.Password)
		if err != nil {
			http.Error(w, "Could not register user", http.StatusInternalServerError)
			return
		}

		// Genera un token para el usuario recién registrado.
		token, err := authService.GenerateToken(newUser.Username, newUser.Password)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}

		// Prepara la respuesta con el token generado y la envía en formato JSON.
		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// authenticateUserHandler es el manejador HTTP para el endpoint de autenticación de usuarios.
// Recibe un repositorio de usuarios (userRepo) y un servicio de autenticación (authService).
func authenticateUserHandler(userRepo user.UserRepository, authService *authentication.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials user.User
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Busca al usuario por su nombre de usuario utilizando el repositorio de usuarios.
		user, err := userRepo.FindByUsername(credentials.Username)
		if err != nil || user.Password != credentials.Password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// El usuario ha sido autenticado correctamente, genera un token para él.
		token, err := authService.GenerateToken(credentials.Username, credentials.Password)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}

		// Prepara la respuesta con el token generado y la envía en formato JSON.
		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
