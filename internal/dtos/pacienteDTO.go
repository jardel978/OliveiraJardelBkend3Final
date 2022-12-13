package dtos

import (
	"gorm.io/gorm"
	"time"
)

type PacienteRequestBody struct {
	Nome                  string    `json:"nome"`
	Sobrenome             string    `json:"sobrenome"`
	DataNascimento        time.Time `json:"data_nascimento" `
	RG                    string    `json:"rg"`
	EnderecoRequestBody   `json:"endereco"`
	ProntuarioRequestBody `json:"prontuario"`
	//domain.Endereco   `json:"endereco"`
	//domain.Prontuario `json:"prontuario"`
	//DataCadastro   time.Time `json:"data_cadastro"`
}

type PacienteResponseBody struct {
	ID                     uint           `json:"id"`
	CreatedAt              time.Time      `json:"created_at,omitempty"`
	UpdatedAt              time.Time      `json:"updated_at,omitempty"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome                   string         `json:"nome"`
	Sobrenome              string         `json:"sobrenome"`
	NomeCompleto           string         `json:"nome_completo"`
	RG                     string         `json:"rg"`
	DataNascimento         time.Time      `json:"data_nascimento"`
	EnderecoResponseBody   `json:"endereco"`
	ProntuarioResponseBody `json:"prontuario"`
	//DataCadastro   time.Time `json:"data_cadastro"`
}
