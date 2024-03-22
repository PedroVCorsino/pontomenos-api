package main

import (
	"log"
	"pontomenos-api/controllers"
	"pontomenos-api/infrastructure/database"
	"pontomenos-api/infrastructure/rabbitMQ"
	"pontomenos-api/infrastructure/repositories"
	"pontomenos-api/models"
	"pontomenos-api/queue/listener"
	"pontomenos-api/queue/sender"
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
	//PostgreSQL
    db := database.ConectarBancoDeDados()
    db.AutoMigrate(&models.Usuario{}, &models.RegistroPonto{})

	//Utils
	authUtils := utils.NewAutenticacaoUtils()

	//Repos
	usuarioRepo := repositories.NewUsuarioRepository(db)
	registroPontoRepo := repositories.NewRegistroPontoRepository(db)

	//Services
	registroPontoService := services.NewRegistroPontoService(usuarioRepo, registroPontoRepo)
    usuarioService := services.NewUsuarioService(usuarioRepo, authUtils)

	//Controllers
    usuarioController := controllers.NewUsuarioController(usuarioService)
	registroPontoController := controllers.NewRegistroPontoController(registroPontoService)

	//RabbitMQ
	rabbitMQConn := rabbitMQ.ConectarRabbitMQ()
    defer rabbitMQConn.Close()
	pontoSender := sender.NewPontoSender(rabbitMQConn)

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir canal: %s", err)
	}
	defer rabbitMQChannel.Close()
	pontoListener := listener.NewPontoListener(rabbitMQChannel, registroPontoService)
	go pontoListener.Listen("registroPontoEvento")
	
	//Rotas
    router := routes.SetupRouter(usuarioController, registroPontoController, usuarioService, pontoSender)
    router.Run() //  ex caso eu queira mudar a porta router.Run("8080")
}
