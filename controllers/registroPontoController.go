package controllers

import (
	"log"
	"net/http"
	"pontomenos-api/services"
	"strconv"
	"time"

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
    resposta, err := rpc.Service.FindRegistroById(id)
    if err != nil {
        if err == services.ErrRegistroNaoEncontrado {
            c.JSON(http.StatusNotFound, gin.H{"error": "registro de ponto não encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    // Formata a resposta
    registroPonto := struct {
        ID        uint      `json:"id"`
        UsuarioID uint      `json:"usuario_id"`
        DataHora  string    `json:"data_hora"` 
        TipoPonto string    `json:"tipo_ponto"` 
    }{
        ID:        resposta.ID,
        UsuarioID: resposta.UsuarioID,
        DataHora:  resposta.DataHora.Format("02/01/2006 15:04"), 
        TipoPonto: resposta.TipoPonto.String(), 
    }

    c.JSON(http.StatusOK, registroPonto)
}

// VisualizarRegistrosPorData @Summary Visualiza registros de ponto por data
// @Security Bearer
// @Description Retorna registros de ponto de um usuário para uma data específica, incluindo o total de horas trabalhadas
// @Tags registros
// @Accept json
// @Produce json
// @Param usuario_id path uint true "ID do Usuário"
// @Param data query string true "Data dos Registros de Ponto (formato: YYYY-MM-DD)"
// @Success 200 {array} RegistroPontoResponse "Uma lista de registros de ponto com total de horas trabalhadas"
// @Router /registros/visualizar [get]
func (rpc *RegistroPontoController) VisualizarRegistrosPorData(c *gin.Context) {
    userID, ok := c.Get("userID")
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "userID não encontrado"})
        return
    }

    userIDStr, ok := userID.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "userID deve ser uma string"})
        return
    }

    usuarioID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
        return
    }

    dataStr := c.Query("data")
    data, err := time.Parse("2006-01-02", dataStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido"})
        return
    }

    registros, totalHoras, err := rpc.Service.VisualizarRegistrosPorData(uint(usuarioID), data) 
    if err != nil {
        if err == services.ErrRegistroNaoEncontrado {
            c.JSON(http.StatusNotFound, gin.H{"error": "Registros de ponto não encontrados para a data"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    resposta := make([]RegistroPontoResponse, 0, len(registros))
    for _, reg := range registros {
        resposta = append(resposta, RegistroPontoResponse{
            ID:        reg.ID,
            UsuarioID: reg.UsuarioID,
            DataHora:  reg.DataHora.Format("02/01/2006 15:04"),
            TipoPonto: reg.TipoPonto.String(),
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "registros":   resposta,
        "total_horas": totalHoras.String(),
    })
}

// EnviarEspelhoMensalPorEmail @Summary Enviar espelho de ponto mensal por e-mail
// @Security Bearer
// @Description Gera e envia por e-mail o espelho de ponto mensal para o usuário especificado.
// @Tags registros
// @Accept  json
// @Produce  json
// @Param   usuarioID  query  string  true  "ID do Usuário"
// @Param   email      query  string  true  "E-mail do Destinatário"
// @Param   mes        query  int     true  "Mês do Relatório"
// @Param   ano        query  int     true  "Ano do Relatório"
// @Success 200 {object} map[string]interface{} "message: Relatório enviado com sucesso"
// @Failure 400 {object} map[string]string "error: Descrição do erro"
// @Failure 500 {object} map[string]string "error: Erro ao enviar e-mail"
// @Router /registros/enviar-espelho [get]
func (rpc *RegistroPontoController) EnviarEspelhoMensalPorEmail(c *gin.Context) {
	log.Println("A")
    usuarioID, err := strconv.ParseUint(c.Query("usuarioID"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
        return
    }
    log.Println("B")
    email := c.Query("email")
    if email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "E-mail destinatário é obrigatório"})
        return
    }
	log.Println("C")
    mes, err := strconv.Atoi(c.Query("mes"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Mês inválido"})
        return
    }
    log.Println("D")
    ano, err := strconv.Atoi(c.Query("ano"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Ano inválido"})
        return
    }
	log.Println("E")
    relatorio, err := rpc.Service.GerarEspelhoPontoMensal(uint(usuarioID), ano, mes)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar relatório"})
        return
    }
	log.Println("F")
    err = rpc.Service.EnviarEmail(relatorio, email, mes, ano)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar e-mail"})
        return
    }
	log.Println("G")
    c.JSON(http.StatusOK, gin.H{"message": "Relatório enviado com sucesso"})
}


