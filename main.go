package main

import (
	"authentication-service/api"
	"authentication-service/authentication"
	"authentication-service/database"
	"authentication-service/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar la base de datos
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Crear el repositorio de usuarios usando la base de datos
	userRepository := &user.PostgresUserRepository{DB: db}

	// Inicializar el servicio de autenticación con el repositorio de usuarios
	authService := authentication.NewAuthService(userRepository)

	// Crear un nuevo enrutador utilizando el paquete "gorilla/mux"
	router := mux.NewRouter()

	// Configurar las rutas del API utilizando el paquete "api"
	// Se pasan el repositorio de usuarios y el servicio de autenticación a las funciones de los handlers.
	api.SetupRoutes(router, userRepository, authService)

	// Imprimir un mensaje indicando que el servidor está iniciando
	println("Starting server on port 3000")

	// Iniciar el servidor HTTP y escuchar en el puerto 3000
	http.ListenAndServe(":3000", router)
}
