package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"LastProject/internal/database"
	"LastProject/internal/model"
)
type UserLogic struct {
	DB *gorm.DB
}

// @Summary Create a new user
// @Description Create a new user in the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param   UserId query int true "UserId"
// @Param   Username query string true "Username"
// @Param   Password query string true "Password"
// @Success 201 {object} model.User
// @Failure 400 "enter correct UserId"
// @Failure 404 "user was not created"
// @Router /users [post]
func (ul UserLogic) CreateUser(c echo.Context) error {
	UserId, err := strconv.Atoi(c.QueryParam("UserId"))
	if err != nil {
		return echo.ErrBadRequest
	}

	Username := c.QueryParam("Username")
	Password, err := GenerateHash(c.QueryParam("Password"))
	if err != nil {
		return err
	}

	user := model.User{UserId: uint(UserId), Username: Username, Password: Password}
	database.AddUserToDataBase(ul.DB, &user)

	return c.JSON(http.StatusCreated, user)
}

// @Summary Get a user by ID
// @Description Get details of a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 401 "Unauthorized"
// @Failure 404 "user not found"
// @Router /users [get]
// @Security BearerAuth
func (ul UserLogic) GetUser(c echo.Context) error {
	id := GetUserIdByToken(c)

	var user model.User
	
	if result := ul.DB.Preload("Orders").First(&user, id); result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, user)
}

// @Summary Delete a user by ID
// @Description Deleting user from db according to ID
// @Tags users
// @Accept json
// @Success 200 "User deleted"
// @Failure 401 "Unauthorized"
// @Failure 404 "User not found"
// @Router /users/{UserId} [delete]
// @Security BearerAuth
func (ul UserLogic) DeleteUser (c echo.Context) error {
	id := GetUserIdByToken(c)

	if err := ul.DB.Where("user_id = ?", id).Delete(&model.Order{}).Error; err != nil {
        return err
    }

	if err := ul.DB.Delete(&model.User{}, id).Error; err != nil {
        return err
    }

	return nil
}