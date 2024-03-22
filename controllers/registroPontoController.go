package controllers

import (
	"net/http"
	"pontomenos-api/services"

	"github.com/gin-gonic/gin"
)

type RegistroPontoController struct {
    Service *services.RegistroPontoService
}

func NewRegistroPontoController(service *services.RegistroPontoService) *RegistroPontoController {
    return &RegistroPontoController{
        Service: service,
    }
}

// FindRegistroPontoById @Summary Busca um registro de ponto pelo ID
// @Security Bearer
// @Description Retorna um registro de ponto dado seu ID
// @Tags registros
// @Accept json
// @Produce json
// @Param id path string true "ID do Registro de Ponto"
// @Success 200 {object} models.RegistroPonto
// @Router /registros/{id} [get]
func (rpc *RegistroPontoController) FindRegistroPontoById(c *gin.Context) {
    id := c.Param("id")
    registroPonto, err := rpc.Service.FindRegistroById(id)
    if err != nil {
        if err == services.ErrRegistroNaoEncontrado {
            c.JSON(http.StatusNotFound, gin.H{"error": "registro de ponto não encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, registroPonto)
}

