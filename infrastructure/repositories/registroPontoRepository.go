package repositories

import (
	"pontomenos-api/models"
	"time"

	"gorm.io/gorm"
)

type RegistroPontoRepository struct {
    db *gorm.DB
}

func NewRegistroPontoRepository(db *gorm.DB) *RegistroPontoRepository {
    return &RegistroPontoRepository{db: db}
}

func (r *RegistroPontoRepository) Create(registro *models.RegistroPonto) (*models.RegistroPonto, error) {
    result := r.db.Create(&registro)
    return registro, result.Error
}

func (r *RegistroPontoRepository) FindById(id string) (*models.RegistroPonto, error) {
    var registro models.RegistroPonto
    result := r.db.First(&registro, "id = ?", id)
    return &registro, result.Error
}

func (r *RegistroPontoRepository) Update(id string, updateData map[string]interface{}) error {
    result := r.db.Model(&models.RegistroPonto{}).Where("id = ?", id).Updates(updateData)
    return result.Error
}

func (r *RegistroPontoRepository) Delete(id string) error {
    result := r.db.Delete(&models.RegistroPonto{}, id)
    return result.Error
}

func (repo *RegistroPontoRepository) FindNextEvent(usuarioID uint, dataHora time.Time) (models.TipoPonto, error) {
    var registros []models.RegistroPonto
    inicioDia := time.Date(dataHora.Year(), dataHora.Month(), dataHora.Day(), 0, 0, 0, 0, dataHora.Location())
    fimDia := inicioDia.Add(24 * time.Hour)

    // Buscar todos os registros do usuário na data fornecida
    err := repo.db.Where("usuario_id = ? AND data_hora >= ? AND data_hora < ?", usuarioID, inicioDia, fimDia).Order("data_hora desc").Find(&registros).Error
    if err != nil {
        return 0, err
    }

    // Determinar o próximo tipo de evento baseado no último registro
    if len(registros) == 0 {
        // Se não houver registros para o dia, é a entrada
        return models.Entrada, nil
    }

    ultimoRegistro := registros[0]
    switch ultimoRegistro.TipoPonto {
    case models.Entrada:
        return models.SaidaIntervalo, nil
    case models.SaidaIntervalo:
        return models.EntradaIntervalo, nil
    case models.EntradaIntervalo:
        return models.Saida, nil
    case models.Saida:
        // Se o último registro foi uma saída, então o limite diário foi atingido
        return models.MaxLimit, nil 
    default:
        return 0, gorm.ErrInvalidData 
    }
}

func (r *RegistroPontoRepository) BuscarRegistrosPorData(usuarioID uint, data time.Time) ([]models.RegistroPonto, error) {
    var registros []models.RegistroPonto
    
    inicioDia := time.Date(data.Year(), data.Month(), data.Day(), 0, 0, 0, 0, data.Location())
    fimDia := inicioDia.AddDate(0, 0, 1)

    err := r.db.Where("usuario_id = ? AND data_hora >= ? AND data_hora < ?", usuarioID, inicioDia, fimDia).
        Order("data_hora ASC"). 
        Find(&registros).Error

    return registros, err
}

func (repo *RegistroPontoRepository) BuscarRegistrosPorPeriodo(usuarioID uint, inicio, fim time.Time) ([]models.RegistroPonto, error) {
    var registros []models.RegistroPonto
    err := repo.db.
        Where("usuario_id = ? AND data_hora >= ? AND data_hora <= ?", usuarioID, inicio, fim).
        Order("data_hora asc").
        Find(&registros).Error
    return registros, err
}