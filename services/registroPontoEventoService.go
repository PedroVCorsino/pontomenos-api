package services

import (
	"pontomenos-api/queue/sender"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type RegistroPontoEventoService struct {
	pontoSender *sender.PontoSender
}

func NewRegistroPontoEventoService(pontoSender *sender.PontoSender) *RegistroPontoEventoService {
	return &RegistroPontoEventoService{
		pontoSender: pontoSender,
	}
}

func (rpes *RegistroPontoEventoService) RegistraPonto(c *gin.Context) error {
	// não sei fazer try_catch ne nessa linguagem, to ficando maluco bixo (╯ ͠° ͟ʖ ͡°)╯┻━┻
	// log.Println("Token Recebido na service:", c.GetHeader("Authorization"))
	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("pontomenos-fiap"), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		login, ok := (*claims)["sub"].(string)
		if !ok {
			return err 
		}

		return rpes.pontoSender.EnviaRegistroParaFila(login)
	}

	return err 
}
