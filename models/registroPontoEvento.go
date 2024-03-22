package models

import "time"

// representa o registro de ponto antes do mesmo ser gravado no banco.
type RegistroPontoEvento struct {
    DataHora time.Time
    Login    string
}
