package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"pontomenos-api/infrastructure/repositories"
	"pontomenos-api/models"
	"time"

	"gopkg.in/gomail.v2"
)

var (
	ErrRegistroNaoEncontrado = errors.New("registro não encontrado")
	ErrMaxLimitAlcancado     = errors.New("máximo de registros diários alcançado")
)

type RegistroPontoService struct {
	usuarioRepo         *repositories.UsuarioRepository
	registroPontoRepo   *repositories.RegistroPontoRepository
}

func NewRegistroPontoService(usuarioRepo *repositories.UsuarioRepository, registroPontoRepo *repositories.RegistroPontoRepository) *RegistroPontoService {
	return &RegistroPontoService{
		usuarioRepo:       usuarioRepo,
		registroPontoRepo: registroPontoRepo,
	}
}

func (rps *RegistroPontoService) ProcessarMensagem(mensagem []byte) {
	var evento models.RegistroPontoEvento
	if err := json.Unmarshal(mensagem, &evento); err != nil {
		log.Printf("Erro ao decodificar a mensagem: %s", err)
		return
	}

	usuario, err := rps.usuarioRepo.FindByLogin(evento.Login)
	if err != nil {
		log.Printf("Usuário não encontrado: %s", err)
		return
	}

	tipoPonto, err := rps.registroPontoRepo.FindNextEvent(usuario.ID, evento.DataHora)
	if err != nil {
		log.Printf("Erro ao buscar o próximo evento: %s", err)
		return
	}

	if tipoPonto == models.MaxLimit {
		log.Println("Limite máximo de registros atingido para o usuário no dia.")
		return
	}

	registro := models.RegistroPonto{
		UsuarioID: usuario.ID,
		DataHora:  evento.DataHora,
		TipoPonto: tipoPonto,
	}

	if _, err := rps.registroPontoRepo.Create(&registro); err != nil {
		log.Printf("Erro ao criar registro de ponto: %s", err)
		return
	}

	log.Println("Registro de ponto criado com sucesso.")
}

func (rps *RegistroPontoService) FindRegistroById(id string) (*models.RegistroPonto, error) {
	registro, err := rps.registroPontoRepo.FindById(id)
	if err != nil {
		return nil, ErrRegistroNaoEncontrado
	}
	return registro, nil
}

func (rps *RegistroPontoService) UpdateRegistro(id string, updateData map[string]interface{}) error {
	return rps.registroPontoRepo.Update(id, updateData)
}

func (rps *RegistroPontoService) DeleteRegistro(id string) error {
	return rps.registroPontoRepo.Delete(id)
}

func (rps *RegistroPontoService) VisualizarRegistrosPorData(usuarioID uint, data time.Time) ([]models.RegistroPonto, time.Duration, error) {
    registros, err := rps.registroPontoRepo.BuscarRegistrosPorData(usuarioID, data)
    if err != nil {
        return nil, 0, ErrRegistroNaoEncontrado
    }
    
    var totalHoras time.Duration
    for i, registro := range registros {
        
        if i%2 == 1 && i > 0 { // Calcula a diferença de tempo entre entrada e saída (intervalo ou fim do dia)
            totalHoras += registro.DataHora.Sub(registros[i-1].DataHora)
        }
    }

    return registros, totalHoras, nil
}

func (rps *RegistroPontoService) GerarEspelhoPontoMensal(usuarioID uint, ano, mes int) (string, error) {
    inicioMes := time.Date(ano, time.Month(mes), 1, 0, 0, 0, 0, time.UTC)
    fimMes := inicioMes.AddDate(0, 1, 0).Add(-time.Nanosecond) 

    registros, err := rps.registroPontoRepo.BuscarRegistrosPorPeriodo(usuarioID, inicioMes, fimMes)
    if err != nil {
        return "", err
    }
	
    var totalHoras time.Duration
    var entrada, saida time.Time
    relatorio := "Data, Tipo de Ponto, Hora\n"
    
    for _, registro := range registros {
        relatorio += fmt.Sprintf("%s, %s, %s\n", registro.DataHora.Format("02/01/2006"), registro.TipoPonto.String(), registro.DataHora.Format("15:04"))
        
        switch registro.TipoPonto {
        case models.Entrada, models.EntradaIntervalo:
            entrada = registro.DataHora
        case models.SaidaIntervalo, models.Saida:
            saida = registro.DataHora
            totalHoras += saida.Sub(entrada)
        }
    }
	
    relatorio += fmt.Sprintf("\nTotal de Horas Trabalhadas no Mês: %v horas", totalHoras.Hours())

    return relatorio, nil
}

func (rps *RegistroPontoService) EnviarEmail(relatorio string, destinatario string, mes int, ano int) error {
    d := gomail.NewDialer("smtp.gmail.com", 587, "pedrovcorsino@gmail.com", "mtiu cbry dviw pwny")

    m := gomail.NewMessage()
    m.SetHeader("From", "pedrovcorsino@gmail.com")
    m.SetHeader("To", destinatario)
    m.SetHeader("Subject", fmt.Sprintf("Espelho de Ponto %d/%d", mes, ano))
    m.SetBody("text/plain", relatorio)
	log.Println("b")
    // Enviar o e-mail
    if err := d.DialAndSend(m); err != nil {
		log.Println(err)
        return err
    }
	
    return nil
}
