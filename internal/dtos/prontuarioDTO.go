package dtos

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
	"time"
)

type EvolucaoTratamentoRequestBody struct {
	ProntuarioID uint   `json:"prontuario"`
	Descricao    string `json:"descricao"`
}

type EvolucaoTratamentoResponseBody struct {
	ID                uint           `json:"id"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty"`
	domain.Prontuario `json:"prontuario"`
	Descricao         string `json:"descricao,omitempty"`
}

type PlanoTratamentoRequestBody struct {
	ProntuarioID uint   `json:"prontuario"`
	Descricao    string `json:"descricao"`
}

type PlanoTratamentoResponseBody struct {
	ID                uint           `json:"id"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty"`
	domain.Prontuario `json:"prontuario"`
	Descricao         string `json:"descricao,omitempty"`
}

type ProntuarioRequestBody struct {
	Consultas           []domain.Consulta           `json:"consultas"`
	PlanosTratamento    []domain.PlanoTratamento    `json:"planos_tratamento"`
	EvolucoesTratamento []domain.EvolucaoTratamento `json:"evolucoes_tratamento"`
}

type ProntuarioResponseBody struct {
	ID                  uint                        `json:"id"`
	CreatedAt           time.Time                   `json:"created_at,omitempty"`
	UpdatedAt           time.Time                   `json:"updated_at,omitempty"`
	DeletedAt           gorm.DeletedAt              `json:"deleted_at,omitempty"`
	Consultas           []domain.Consulta           `json:"consultas,omitempty"`
	PlanosTratamento    []domain.PlanoTratamento    `json:"planos_tratamento,omitempty"`
	EvolucoesTratamento []domain.EvolucaoTratamento `json:"evolucoes_tratamento,omitempty"`
}
