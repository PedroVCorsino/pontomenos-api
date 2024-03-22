package routes

import (
	"pontomenos-api/controllers"
	middleware "pontomenos-api/middlewares"

	"github.com/gin-gonic/gin"
)

func UsuarioRoutes(router *gin.Engine, usuarioController *controllers.UsuarioController) {
    usuario := router.Group("/usuarios")

    // A rota de criar usuário permanece pública
    usuario.POST("/", usuarioController.CreateUsuario)

    // Criando um subgrupo para rotas que requerem autenticação
    protected := usuario.Group("/")
    protected.Use(middleware.AuthorizeJWT()) // Aplica o middleware JWT
    {
        protected.GET("/:id", usuarioController.FindUsuarioById)
        protected.PATCH("/:id", usuarioController.UpdateUsuario)
        protected.DELETE("/:id", usuarioController.DeleteUsuario)
    }
}
