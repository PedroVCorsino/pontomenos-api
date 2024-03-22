package services

import (
    "errors"
    "pontomenos-api/infrastructure/repositories"
    "pontomenos-api/models"
    "pontomenos-api/utils"
)

var ErrUsuarioNaoEncontrado = errors.New("usuário não encontrado")

type UsuarioService struct {
    repo       *repositories.UsuarioRepository
    authUtils  *utils.AutenticacaoUtils 
}

func NewUsuarioService(repo *repositories.UsuarioRepository, authUtils *utils.AutenticacaoUtils) *UsuarioService {
    return &UsuarioService{
        repo:      repo,
        authUtils: authUtils,
    }
}

func (us *UsuarioService) CreateUsuario(usuario *models.Usuario) (*models.Usuario, error) {
    senhaHash, err := us.authUtils.HashSenha(usuario.Senha)
    if err != nil {
        return nil, err
    }
    usuario.Senha = senhaHash
    return us.repo.Create(usuario)
}

func (us *UsuarioService) FindUsuarioById(id string) (*models.Usuario, error) {
    usuario, err := us.repo.FindById(id)
    if err != nil {
        return nil, ErrUsuarioNaoEncontrado
    }
    return usuario, nil
}

func (us *UsuarioService) UpdateUsuario(id string, updateData map[string]interface{}) error {
    if senha, ok := updateData["senha"].(string); ok && senha != "" {
        senhaHash, err := us.authUtils.HashSenha(senha)
        if err != nil {
            return err
        }
        updateData["senha"] = senhaHash
    }

    return us.repo.Update(id, updateData)
}

func (us *UsuarioService) DeleteUsuario(id string) error {
    return us.repo.Delete(id)
}

func (us *UsuarioService) FindByLogin(login string) (*models.Usuario, error) {
    return us.repo.FindByLogin(login)
}
