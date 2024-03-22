package routes

import (
	"pontomenos-api/controllers"
	_ "pontomenos-api/docs"
	"pontomenos-api/queue"
	"pontomenos-api/services"
	"pontomenos-api/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(usuarioController *controllers.UsuarioController, 
    usuarioService *services.UsuarioService, 
    pontoSender *queue.PontoSender) *gin.Engine {
router := gin.Default()

    authUtils := utils.NewAutenticacaoUtils()

    // Configuração do Login
    loginService := services.NewLoginService(usuarioService, authUtils)
    loginController := controllers.NewLoginController(loginService)
    router.POST("/auth", loginController.Autenticar)

    // Configuração do Registro de Ponto
    registroPontoService := services.NewRegistroPontoEventoService(pontoSender)
    registroPontoController := controllers.NewRegistroPontoEventoController(registroPontoService)
    router.POST("/ponto", registroPontoController.RegistraPonto)
    
    // Configuração do Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Rotas de Usuário
    UsuarioRoutes(router, usuarioController)
    
    return router
}
