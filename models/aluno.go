package models

type Aluno struct {
	Nome string `json:"nomes"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno
