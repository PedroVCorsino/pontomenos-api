package routes

import (
	"pontomenos-api/controllers"
	middleware "pontomenos-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegistroPontoRoutes(router *gin.Engine, registroPontoController *controllers.RegistroPontoController) {
    registro := router.Group("/registros")
    registro.Use(middleware.AuthorizeJWT()) 
    {
        registro.GET("/:id", registroPontoController.FindRegistroPontoById)
        // registro.PATCH("/:id", registroPontoController.UpdateRegistroPonto)  
        // registro.DELETE("/:id", registroPontoController.DeleteRegistroPonto) 
    }
}
