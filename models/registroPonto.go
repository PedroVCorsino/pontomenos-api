package models

import (
	"time"
)

// TipoPonto define um tipo personalizado para representar os tipos de ponto.
type TipoPonto int

// Constantes que representam os tipos de ponto válidos. 
// To tentando simualar uma enum. (ง ͠° ͟ل͜ ͡°)ง
const (
	Entrada TipoPonto = iota // iota promove a enumeração automática, começando de 0
	SaidaIntervalo
	EntradaIntervalo
	Saida
	MaxLimit
)

// RegistroPonto representa um registro de ponto de um usuário.
type RegistroPonto struct {
	ID        uint      `gorm:"primaryKey"`
	UsuarioID uint      `json:"usuario_id"`
	DataHora  time.Time `json:"data_hora"`
	TipoPonto TipoPonto `json:"tipo_ponto"`
}
