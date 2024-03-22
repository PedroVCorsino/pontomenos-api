package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files" 
    "github.com/swaggo/gin-swagger" 
    _ "pontomenos-api/docs" 
    "pontomenos-api/controllers"
    "pontomenos-api/services"
    "pontomenos-api/utils" 
)

func SetupRouter(usuarioController *controllers.UsuarioController, usuarioService *services.UsuarioService) *gin.Engine {
    router := gin.Default()

    authUtils := utils.NewAutenticacaoUtils()

    loginService := services.NewLoginService(usuarioService, authUtils)
    loginController := controllers.NewLoginController(loginService)

    UsuarioRoutes(router, usuarioController)

    router.POST("/auth", loginController.Autenticar)

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return router
}
