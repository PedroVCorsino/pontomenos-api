package utils

import (
    "golang.org/x/crypto/bcrypt"
)

type AutenticacaoUtils struct{}

func NewAutenticacaoUtils() *AutenticacaoUtils {
    return &AutenticacaoUtils{}
}

func (au *AutenticacaoUtils) HashSenha(senha string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
    return string(hash), err
}

func (au *AutenticacaoUtils) VerificarSenha(senha string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
    return err == nil
}
