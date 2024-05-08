package database

import (
	"api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDataBase() {
	stringConextion := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConextion))

	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
