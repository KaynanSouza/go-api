package main

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"api-go-gin/routes"
)

func main() {

	database.ConectDataBase()

	models.Alunos = []models.Aluno{
		{Nome: "Kaynan", CPF: "10000000000", RG: "251000000"},
		{Nome: "Murilo", CPF: "20000000000", RG: "252000000"},
		{Nome: "Giovana", CPF: "30000000000", RG: "253000000"},
		{Nome: "Pedro", CPF: "40000000000", RG: "254000000"},
		{Nome: "Felp", CPF: "50000000000", RG: "255000000"},
	}
	routes.HandleRequests()
}
