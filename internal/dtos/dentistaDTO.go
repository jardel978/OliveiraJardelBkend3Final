package dtos

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
	"time"
)

type DentistaRequestBody struct {
	Nome            string `json:"nome" binding:"required"`
	Sobrenome       string `json:"sobrenome" binding:"required"`
	domain.Endereco `json:"endereco"`
	Matricula       string            `json:"matricula" binding:"required"`
	Consultas       []domain.Consulta `json:"consultas"`
}

type DentistaResponseBody struct {
	ID              uint           `json:"id"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome            string         `json:"nome"`
	Sobrenome       string         `json:"sobrenome"`
	NomeCompleto    string         `json:"nome_completo"`
	domain.Endereco `json:"endereco"`
	Matricula       string            `json:"matricula"`
	Consultas       []domain.Consulta `json:"consultas"`
}
