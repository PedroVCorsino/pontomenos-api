package controllers

import (
	"net/http"
	"pontomenos-api/services" // Ajuste para o caminho correto do seu pacote services

	"github.com/gin-gonic/gin"
)

// LoginData representa os dados de login recebidos na requisição
type LoginData struct {
    Login string `json:"login"`
    Senha string `json:"senha"`
}

// LoginController controla as operações de login
type LoginController struct {
    loginService *services.LoginService
}

// NewLoginController cria uma nova instância de LoginController
func NewLoginController(loginService *services.LoginService) *LoginController {
    return &LoginController{
        loginService: loginService,
    }
}

// @Summary Autentica um usuário
// @Description Autentica um usuário com login e senha
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   body body LoginData true "Dados de Login"
// @Success 200 {object} map[string]string "Token JWT"
// @Failure 400 {object} map[string]string "Mensagem de erro para requisição inválida"
// @Failure 401 {object} map[string]string "Mensagem de erro para login ou senha inválidos"
// @Router /auth [post]
func (lc *LoginController) Autenticar(c *gin.Context) {
    var loginData LoginData
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
        return
    }

    token, err := lc.loginService.Autenticar(loginData.Login, loginData.Senha)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Login ou senha inválidos"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
