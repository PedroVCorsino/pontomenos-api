package main

import (
	"log"
	"pontomenos-api/models"
	"pontomenos-api/utils"
	"time"

	"gorm.io/gorm"
)

func gerarRegistrosUteis(db *gorm.DB, usuarioID uint, inicio time.Time, diasUteis int) {
    horarios := []struct {
        Hora  int
        Min   int
        TipoPonto models.TipoPonto
    }{
        {8, 0, models.Entrada},
        {12, 0, models.SaidaIntervalo},
        {13, 0, models.EntradaIntervalo},
        {18, 0, models.Saida},
    }

    adicionados := 0
    for adicionados < diasUteis {
        if inicio.Weekday() != time.Saturday && inicio.Weekday() != time.Sunday {
            for _, h := range horarios {
                registro := models.RegistroPonto{
                    UsuarioID: usuarioID,
                    DataHora:  time.Date(inicio.Year(), inicio.Month(), inicio.Day(), h.Hora, h.Min, 0, 0, inicio.Location()),
                    TipoPonto: h.TipoPonto,
                }
                if err := db.Create(&registro).Error; err != nil {
                    log.Fatalf("Erro ao criar registro de ponto: %v", err)
                }
            }
            adicionados++
        }
        inicio = inicio.AddDate(0, 0, 1)
    }
}

func popularBancoDeDados(db *gorm.DB) {
    senhaHash, err := utils.NewAutenticacaoUtils().HashSenha("mudar123")
    if err != nil {
        log.Fatalf("Erro ao criar hash da senha: %v", err)
    }

    usuario := models.Usuario{Nome: "Pedro Vinicius", Matricula: "000001", Email: "pedroviniciuscorsino@gmail.com", Senha: senhaHash}
    if err := db.Create(&usuario).Error; err != nil {
        log.Fatalf("Erro ao criar usuário: %v", err)
    }

    inicioFevereiro := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
    inicioMarco := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

    gerarRegistrosUteis(db, usuario.ID, inicioFevereiro, 20) // Ajuste o número de dias úteis conforme necessário
    gerarRegistrosUteis(db, usuario.ID, inicioMarco, 15) // Ajuste para os primeiros 15 dias úteis de março
}
