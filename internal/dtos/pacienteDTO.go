package dtos

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
	"time"
)

type PacienteRequestBody struct {
	Nome              string    `json:"nome" binding:"required"`
	Sobrenome         string    `json:"sobrenome" binding:"required"`
	DataNascimento    time.Time `json:"data_nascimento"  binding:"required"`
	domain.Endereco   `json:"endereco"`
	domain.Prontuario `json:"prontuario"`
	//DataCadastro   time.Time `json:"data_cadastro"`
}

type PacienteResponseBody struct {
	ID                uint           `json:"id"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty"`
	Nome              string         `json:"nome"`
	Sobrenome         string         `json:"sobrenome"`
	NomeCompleto      string         `json:"nome_completo"`
	DataNascimento    time.Time      `json:"data_nascimento"`
	domain.Endereco   `json:"endereco"`
	domain.Prontuario `json:"prontuario"`
	//DataCadastro   time.Time `json:"data_cadastro"`
}
