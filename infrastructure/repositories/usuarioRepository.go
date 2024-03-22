// infrastructure/repositories/usuarioRepository.go

package repositories

import (
    "pontomenos-api/models"
    "gorm.io/gorm"
)

type UsuarioRepository struct {
    db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
    return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) (*models.Usuario, error) {
    result := r.db.Create(&usuario)
    return usuario, result.Error
}

func (r *UsuarioRepository) FindById(id string) (*models.Usuario, error) {
    var usuario models.Usuario
    result := r.db.First(&usuario, "id = ?", id)
    return &usuario, result.Error
}

func (r *UsuarioRepository) Update(id string, updateData map[string]interface{}) error {
    resultado := r.db.Model(&models.Usuario{}).Where("id = ?", id).Updates(updateData)
    if resultado.Error != nil {
        return resultado.Error
    }
    return nil
}

func (r *UsuarioRepository) Delete(id string) error {
    result := r.db.Delete(&models.Usuario{}, id)
    return result.Error
}

// repositories/usuarioRepository.go

func (r *UsuarioRepository) FindByLogin(login string) (*models.Usuario, error) {
    var usuario models.Usuario
    if err := r.db.Where("nome = ? OR email = ?", login, login).First(&usuario).Error; err != nil {
        return nil, err
    }
    return &usuario, nil
}
