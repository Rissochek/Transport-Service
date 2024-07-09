package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"LastProject/internal/database"
	"LastProject/internal/handlers"
	_ "LastProject/docs"
)

func InitApp() {
	e := echo.New()
	db, err := database.InitDataBase()
	if err != nil{
		log.Fatal(err)
	}
	
	secretKey := handlers.GetSecretKey()

	ul := handlers.UserLogic{DB:db}
	ol := handlers.OrderLogic{DB:db}
	al := handlers.AuthLogic{DB:db}

	e.GET("/api/v1/login", al.AuthLogin)
    e.POST("/api/v1/users", ul.CreateUser)
	
	r := e.Group("/api/v1")
	config := echojwt.Config{
		SigningKey: []byte(secretKey),
	}
	
	r.Use(echojwt.WithConfig(config))

	r.GET("/users", ul.GetUser)
	r.DELETE("/users/:UserId", ul.DeleteUser)

	r.POST("/orders", ol.CreateOrder)
    r.GET("/orders/:OrderId", ol.GetOrder)
	r.PUT("/orders/:OrderId", ol.UpdateOrder)
	r.DELETE("/orders/:OrderId", ol.DeleteOrder)
    
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    e.Logger.Fatal(e.Start(":8080"))
}