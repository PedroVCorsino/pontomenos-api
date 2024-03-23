package models

// Usuario representa um usuário do sistema
type Usuario struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Nome      string `json:"nome"`
    Matricula string `gorm:"unique" json:"matricula"`
    Email     string `gorm:"unique" json:"email"`
    Senha     string `json:"senha"`
}
