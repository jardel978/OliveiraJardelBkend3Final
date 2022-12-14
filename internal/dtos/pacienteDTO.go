package dtos

import (
	"gorm.io/gorm"
	"time"
)

type PacienteRequestBody struct {
	ID                    uint   `json:"id"`
	Nome                  string `json:"nome"`
	Sobrenome             string `json:"sobrenome"`
	DataNascimento        string `json:"data_nascimento" `
	RG                    string `json:"rg"`
	EnderecoRequestBody   `json:"endereco"`
	ProntuarioRequestBody `json:"prontuario"`
}

type PacienteResponseBody struct {
	ID                     uint           `json:"id,omitempty"`
	CreatedAt              time.Time      `json:"created_at,omitempty"`
	UpdatedAt              time.Time      `json:"updated_at,omitempty"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome                   string         `json:"nome,omitempty"`
	Sobrenome              string         `json:"sobrenome,omitempty"`
	NomeCompleto           string         `json:"nome_completo,omitempty"`
	RG                     string         `json:"rg,omitempty"`
	DataNascimento         time.Time      `json:"data_nascimento,omitempty"`
	EnderecoResponseBody   `json:"endereco,omitempty"`
	ProntuarioResponseBody `json:"prontuario,omitempty"`
}
