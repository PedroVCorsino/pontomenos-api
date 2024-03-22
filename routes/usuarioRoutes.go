package routes

import (
    "github.com/gin-gonic/gin"
    "pontomenos-api/controllers"
)

func UsuarioRoutes(router *gin.Engine, usuarioController *controllers.UsuarioController) {
    usuario := router.Group("/usuarios")
    {
        usuario.POST("/", usuarioController.CreateUsuario)
        usuario.GET("/:id", usuarioController.FindUsuarioById)
        usuario.PATCH("/:id", usuarioController.UpdateUsuario)
        usuario.DELETE("/:id", usuarioController.DeleteUsuario)
    }
}
