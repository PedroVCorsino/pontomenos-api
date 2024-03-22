package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "pontomenos-api/services"
)

type LoginController struct {
    loginService *services.LoginService
}

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
// @Param   login body string true "Login do Usuário"
// @Param   senha body string true "Senha do Usuário"
// @Success 200 {string} string "token JWT"
// @Router /auth [post]
func (lc *LoginController) Autenticar(c *gin.Context) {
    var loginData struct {
        Login string `json:"login"`
        Senha string `json:"senha"`
    }
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
