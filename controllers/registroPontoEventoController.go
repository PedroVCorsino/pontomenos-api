package controllers

import (
	"net/http"
	"pontomenos-api/services"

	"github.com/gin-gonic/gin"
)

type RegistroPontoEventoController struct {
	registroPontoService *services.RegistroPontoEventoService
}

func NewRegistroPontoEventoController(registroPontoService *services.RegistroPontoEventoService) *RegistroPontoEventoController {
	return &RegistroPontoEventoController{
		registroPontoService: registroPontoService,
	}
}

// RegistraPonto é o endpoint para registrar o ponto do usuário.
// @Summary Registra ponto
// @Description Registra o ponto do usuário com o login obtido do JWT
// @Tags ponto
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 204 "Ponto registrado com sucesso"
// @Failure 401 "Não autorizado"
// @Failure 500 "Erro interno do servidor"
// @Router /ponto [post]
func (rp *RegistroPontoEventoController) RegistraPonto(c *gin.Context) {
	if err := rp.registroPontoService.RegistraPonto(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar ponto"})
		return
	}
	c.Status(http.StatusNoContent)
}
