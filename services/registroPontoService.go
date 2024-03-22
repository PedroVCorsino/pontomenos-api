package services

import (
	"encoding/json"
	"errors"
	"log"
	"pontomenos-api/infrastructure/repositories"
	"pontomenos-api/models"
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
