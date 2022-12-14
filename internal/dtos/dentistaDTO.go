package dtos

import (
	"gorm.io/gorm"
	"time"
)

type DentistaRequestBody struct {
	ID                  uint   `json:"id"`
	Nome                string `json:"nome"`
	Sobrenome           string `json:"sobrenome"`
	EnderecoRequestBody `json:"endereco"`
	Matricula           string                `json:"matricula"`
	ClinicaID           uint                  `json:"clinica_id"`
	Consultas           []ConsultaRequestBody `json:"consultas"`
}

type DentistaResponseBody struct {
	ID                   uint           `json:"id,omitempty"`
	CreatedAt            time.Time      `json:"created_at,omitempty"`
	UpdatedAt            time.Time      `json:"updated_at,omitempty"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome                 string         `json:"nome,omitempty"`
	Sobrenome            string         `json:"sobrenome,omitempty"`
	NomeCompleto         string         `json:"nome_completo,omitempty"`
	EnderecoResponseBody `json:"endereco,omitempty"`
	Matricula            string                 `json:"matricula,omitempty"`
	ClinicasID           uint                   `json:"clinica_id,omitempty"`
	Consultas            []ConsultaResponseBody `json:"consultas,omitempty"`
}
