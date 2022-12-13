package dtos

import (
	"gorm.io/gorm"
	"time"
)

type ConsultaRequestBody struct {
	PacienteID uint   `json:"paciente_id"`
	DentistaID uint   `json:"dentista_id"`
	Data       string `json:"data"`
	Descricao  string `json:"descricao"`
}

type ConsultaResponseBody struct {
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
	PacienteID   uint           `json:"paciente_id"`
	DentistaID   uint           `json:"dentista_id"`
	DataConsulta time.Time      `json:"data"`
	Descricao    string         `json:"descricao,omitempty"`
}
