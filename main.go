package main

import (
	"pontomenos-api/controllers"
	"pontomenos-api/infrastructure/database"
	"pontomenos-api/infrastructure/rabbitMQ"
	"pontomenos-api/infrastructure/repositories"
	"pontomenos-api/models"
	"pontomenos-api/queue"
	"pontomenos-api/routes"
	"pontomenos-api/services"
	"pontomenos-api/utils"
)

// @title Pontomenos API
// @description API desenvolvida para o hackthon da FIAP PósTech.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// Não aguento mais reescrever esse arquivo ಠ_ಠ
func main() {
    db := database.ConectarBancoDeDados()
    db.AutoMigrate(&models.Usuario{})

	rabbitMQConn := rabbitMQ.ConectarRabbitMQ()
    defer rabbitMQConn.Close()
	pontoSender := queue.NewPontoSender(rabbitMQConn)

	authUtils := utils.NewAutenticacaoUtils()
    usuarioRepo := repositories.NewUsuarioRepository(db)
    usuarioService := services.NewUsuarioService(usuarioRepo, authUtils)
    usuarioController := controllers.NewUsuarioController(usuarioService)

    router := routes.SetupRouter(usuarioController, usuarioService, pontoSender)
    router.Run()
}
