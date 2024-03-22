package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "pontomenos-api/models"
    "pontomenos-api/services"
)

// UsuarioController estrutura para o controlador do usuário.
type UsuarioController struct {
    Service *services.UsuarioService
}

// NewUsuarioController cria uma nova instância de UsuarioController.
func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
    return &UsuarioController{
        Service: service,
    }
}

// CreateUsuario @Summary Adiciona um novo usuário
// @Description Adiciona um novo usuário com as informações fornecidas
// @Tags usuarios
// @Accept json
// @Produce json
// @Param usuario body models.Usuario true "Informações do Usuário"
// @Success 201 {object} models.Usuario
// @Router /usuarios [post]
func (uc *UsuarioController) CreateUsuario(c *gin.Context) {
    var usuario models.Usuario
    if err := c.ShouldBindJSON(&usuario); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    res, err := uc.Service.CreateUsuario(&usuario)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, res)
}

// FindUsuarioById @Summary Busca um usuário pelo ID
// @Description Retorna um usuário dado seu ID
// @Tags usuarios
// @Accept json
// @Produce json
// @Param id path string true "ID do Usuário"
// @Success 200 {object} models.Usuario
// @Router /usuarios/{id} [get]
func (uc *UsuarioController) FindUsuarioById(c *gin.Context) {
    id := c.Param("id")
    usuario, err := uc.Service.FindUsuarioById(id)
    if err != nil {
        // Verifica se o erro é devido ao usuário não encontrado
        if err == services.ErrUsuarioNaoEncontrado {
            c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
        } else {
            // Para outros tipos de erro, retorna um erro interno do servidor
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    // Se não houver erro, retorna o usuário encontrado
    c.JSON(http.StatusOK, usuario)
}

// @Summary Atualiza informações de um usuário
// @Description Atualiza parcialmente um usuário existente com as informações fornecidas
// @Tags usuarios
// @Accept json
// @Produce json
// @Param id path string true "ID do Usuário"
// @Param body body map[string]interface{} true "Campos do Usuário para Atualizar"
// @Success 200 {object} map[string]interface{} "Usuário atualizado com sucesso"
// @Router /usuarios/{id} [patch]
func (uc *UsuarioController) UpdateUsuario(c *gin.Context) {
    var updateData map[string]interface{}
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := c.Param("id")
    err := uc.Service.UpdateUsuario(id, updateData)
    if err != nil {
        if err == services.ErrUsuarioNaoEncontrado {
            c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
}


// DeleteUsuario @Summary Exclui um usuário
// @Description Exclui um usuário dado seu ID
// @Tags usuarios
// @Accept json
// @Produce json
// @Param id path string true "ID do Usuário"
// @Success 200 {object} map[string]string "message: Usuário excluído com sucesso!"
// @Router /usuarios/{id} [delete]
func (uc *UsuarioController) DeleteUsuario(c *gin.Context) {
    id := c.Param("id")
    err := uc.Service.DeleteUsuario(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Usuário excluído com sucesso!"})
}
