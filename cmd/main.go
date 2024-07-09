package main

import "LastProject/internal/app"

// @title Swagger Example API
// @version 1.0
// @description Это пример сервера Swagger.
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    app.InitApp()
}
