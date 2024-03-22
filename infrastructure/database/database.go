// infrastructure/database/database.go

package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func ConectarBancoDeDados() *gorm.DB {
    dsn := "host=localhost user=admin password=mudar@123 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Panic("Erro ao conectar com o banco de dados", err)
    }
    return db
}
