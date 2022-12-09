package dtos

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
	"time"
)

type ClinicaRequestBody struct {
	Nome            string `json:"nome" binding:"required"`
	CNPJ            string `json:"cnpj" binding:"required"`
	domain.Endereco `json:"endereco"`
	Dentistas       []domain.Dentista `json:"dentistas"`
}

type ClinicaResponseBody struct {
	ID              uint           `json:"id"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome            string         `json:"nome"`
	CNPJ            string         `json:"cnpj"`
	domain.Endereco `json:"endereco"`
	Dentistas       []domain.Dentista `json:"dentistas"`
}
