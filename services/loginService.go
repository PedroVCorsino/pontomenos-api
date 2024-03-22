package services

import (
	"errors"
	"pontomenos-api/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("sua_chave_secreta_super_secreta")

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
    claims := &jwt.StandardClaims{
        Subject:   usuario.Email,
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}
