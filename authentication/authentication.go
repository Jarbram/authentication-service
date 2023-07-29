package authentication

import (
	"authentication-service/user"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Clave secreta para firmar los tokens JWT.
var secretKey = []byte("mkdir123")

// AuthService representa el servicio de autenticación.
type AuthService struct {
	userRepo user.UserRepository
}

// NewAuthService crea una nueva instancia del servicio de autenticación.
func NewAuthService(userRepo user.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// RegisterUser registra un nuevo usuario en el servicio de autenticación.
// Verifica si el usuario ya existe en la base de datos antes de crearlo.
func (a *AuthService) RegisterUser(username, password string) error {
	// Primero, verifica si el usuario ya existe en la base de datos.
	_, err := a.userRepo.FindByUsername(username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Si el usuario no existe, procede a crearlo en la base de datos.
	newUser := &user.User{Username: username, Password: password}
	return a.userRepo.CreateUser(newUser)
}

// GenerateToken genera un token JWT para el usuario con las credenciales proporcionadas.
func (a *AuthService) GenerateToken(username, password string) (string, error) {
	// Busca al usuario por su nombre de usuario utilizando el repositorio de usuarios.
	user, err := a.userRepo.FindByUsername(username)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return "", err
	}

	// Verifica que la contraseña proporcionada coincida con la contraseña del usuario.
	if user.Password != password {
		fmt.Println("Invalid credentials")
		return "", errors.New("invalid credentials")
	}

	// Crea los claims (datos adicionales) que se incluirán en el token JWT.
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas.
	}

	// Crea un nuevo token JWT firmado con los claims y la clave secreta.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
