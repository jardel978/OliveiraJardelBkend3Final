package dtos

import (
	"gorm.io/gorm"
	"time"
)

type ClinicaRequestBody struct {
	ID                  uint
	Nome                string `json:"nome"`
	CNPJ                string `json:"cnpj"`
	EnderecoRequestBody `json:"endereco"`
	Dentistas           []DentistaRequestBody `json:"dentistas"`
}

type ClinicaResponseBody struct {
	ID                   uint           `json:"id"`
	CreatedAt            time.Time      `json:"created_at,omitempty"`
	UpdatedAt            time.Time      `json:"updated_at,omitempty"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome                 string         `json:"nome"`
	CNPJ                 string         `json:"cnpj"`
	EnderecoResponseBody `json:"endereco"`
	Dentistas            []DentistaResponseBody `json:"dentistas"`
}
