package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("pontomenos-fiap")

func AuthorizeJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        const BearerSchema = "Bearer "
        header := c.GetHeader("Authorization")
        if header == "" {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Header de autorização não encontrado"})
            return
        }

        tokenString := header[len(BearerSchema):]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, errors.New("Erro no método de assinatura do JWT")
            }
            return jwtKey, nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token JWT inválido"})
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("userID", claims["sub"])
        } else {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token JWT inválido ou expirado"})
            return
        }
    }
}
