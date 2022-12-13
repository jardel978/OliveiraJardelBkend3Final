package dtos

import (
	"gorm.io/gorm"
	"time"
)

type DentistaRequestBody struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	//domain.Endereco `json:"endereco"`
	EnderecoRequestBody `json:"endereco_id"`
	Matricula           string                `json:"matricula"`
	Consultas           []ConsultaRequestBody `json:"consultas"`
	Clinicas            []ClinicaRequestBody  `json:"clinica_id"`
}

type DentistaResponseBody struct {
	ID                   uint           `json:"id"`
	CreatedAt            time.Time      `json:"created_at,omitempty"`
	UpdatedAt            time.Time      `json:"updated_at,omitempty"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome                 string         `json:"nome"`
	Sobrenome            string         `json:"sobrenome"`
	NomeCompleto         string         `json:"nome_completo"`
	EnderecoResponseBody `json:"endereco"`
	Matricula            string                 `json:"matricula"`
	Consultas            []ConsultaResponseBody `json:"consultas"`
	Clinicas             []ClinicaResponseBody  `json:"clinicas"`
}
