package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nomes"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno
