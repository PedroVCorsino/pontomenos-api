package routes

import (
	"pontomenos-api/controllers"
	_ "pontomenos-api/docs"
	"pontomenos-api/queue/sender"
	"pontomenos-api/services"
	"pontomenos-api/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(usuarioController *controllers.UsuarioController,
    registroPontoController *controllers.RegistroPontoController, 
    usuarioService *services.UsuarioService, 
    pontoSender *sender.PontoSender) *gin.Engine {
router := gin.Default()

    authUtils := utils.NewAutenticacaoUtils()

    // Configuração do Login
    loginService := services.NewLoginService(usuarioService, authUtils)
    loginController := controllers.NewLoginController(loginService)
    router.POST("/auth", loginController.Autenticar)

    // Configuração do Registro de Ponto
    registroPontoService := services.NewRegistroPontoEventoService(pontoSender)
    registroPontoEventoController := controllers.NewRegistroPontoEventoController(registroPontoService)
    router.POST("/ponto", registroPontoEventoController.RegistraPonto)
    
    // Configuração do Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Rotas de Usuário
    UsuarioRoutes(router, usuarioController)
    RegistroPontoRoutes(router, registroPontoController)
    
    return router
}
