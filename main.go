package main

import (
    "pontomenos-api/controllers"
    "pontomenos-api/infrastructure/database"
    "pontomenos-api/models"
    "pontomenos-api/infrastructure/repositories"
    "pontomenos-api/routes"
    "pontomenos-api/services"
	"pontomenos-api/utils" /

// @title Pontomenos API
// @description API desenvolvida para o hackthon da FIAP PósTech.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
    db := database.ConectarBancoDeDados()
    db.AutoMigrate(&models.Usuario{})

	authUtils := utils.NewAutenticacaoUtils()
    usuarioRepo := repositories.NewUsuarioRepository(db)
    usuarioService := services.NewUsuarioService(usuarioRepo, authUtils)
    usuarioController := controllers.NewUsuarioController(usuarioService)

    router := routes.SetupRouter(usuarioController, usuarioService)
    router.Run()
}
