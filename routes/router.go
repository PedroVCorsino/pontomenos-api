package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files" // swagger embed files
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    _ "pontomenos-api/docs" // O pacote docs é gerado pelo Swaggo, ajuste o caminho conforme necessário
    "pontomenos-api/controllers"
    "pontomenos-api/services"
    "pontomenos-api/utils" // Importe utils para acessar AutenticacaoUtils
)

// SetupRouter configura e retorna um novo roteador Gin.
func SetupRouter(usuarioController *controllers.UsuarioController, usuarioService *services.UsuarioService) *gin.Engine {
    router := gin.Default()

    // Cria uma instância de AutenticacaoUtils
    authUtils := utils.NewAutenticacaoUtils()

    // Assumindo que usuarioService já foi criado e injetado onde SetupRouter é chamado
    loginService := services.NewLoginService(usuarioService, authUtils)
    loginController := controllers.NewLoginController(loginService)

    // Configura as rotas do usuário
    UsuarioRoutes(router, usuarioController)

    // Rota para autenticação
    router.POST("/auth", loginController.Autenticar)

    // Configuração do Swagger UI
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return router
}
