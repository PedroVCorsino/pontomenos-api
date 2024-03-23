package services

import (
	"errors"
	"pontomenos-api/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("pontomenos-fiap")

type LoginService struct {
    userService *UsuarioService
    authUtils   *utils.AutenticacaoUtils
}

func NewLoginService(userService *UsuarioService, authUtils *utils.AutenticacaoUtils) *LoginService {
    return &LoginService{
        userService: userService,
        authUtils:   authUtils,
    }
}

func (ls *LoginService) Autenticar(login, senha string) (string, error) {
    usuario, err := ls.userService.FindByLogin(login)
    if err != nil {
        return "", err
    }

    if !ls.authUtils.VerificarSenha(senha, usuario.Senha) {
        return "", errors.New("login ou senha inválidos")
    }

    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &jwt.RegisteredClaims{
        Subject:   usuario.Matricula,
        ExpiresAt: jwt.NewNumericDate(expirationTime),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}
