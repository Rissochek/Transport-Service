package handlers

import (
	"net/http"
	"time"
	"os"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"LastProject/internal/model"
)

type AuthLogic struct {
	DB *gorm.DB
}

type jwtCustomClaims struct {
	UserId uint
	Username string
	jwt.RegisteredClaims
}

// @Summary Auth in system
// @Description Auth in system with username and password
// @Tags auth
// @Accept  json
// @Produce plain
// @Param   username query string true "Username"
// @Param   password query string true "Password"
// @Success 200 {string} string "JWT token"
// @Failure 401 "Uncorrect password"
// @Failure 404 "User not found"
// @Router /login [get]
func (al AuthLogic) AuthLogin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	var user model.User 

	secretKey := GetSecretKey()

	if err := al.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	if status := CompareHashAndPassword(password, user.Password); !status{
		return echo.ErrUnauthorized
	}
	
	claims := &jwtCustomClaims{
		user.UserId,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return err
	}
	BearerToken := "Bearer "+ t
	return c.String(http.StatusOK, BearerToken)
}

func GetUserIdByToken (c echo.Context) uint {
	VerifiedUser := c.Get("user").(*jwt.Token)
	claims := VerifiedUser.Claims.(jwt.MapClaims)
	idFloat := claims["UserId"].(float64)
	id := uint(idFloat)

	return id
}

func GetSecretKey() string{
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatalf("SECRET_KEY не установлен в .env файле")
	}
	return secretKey
}