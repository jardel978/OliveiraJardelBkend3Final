package dtos

import (
	"gorm.io/gorm"
	"time"
)

type ConsultaRequestBody struct {
	ID           uint   `json:"id"`
	PacienteID   uint   `json:"paciente_id"`
	DentistaID   uint   `json:"dentista_id"`
	DataConsulta string `json:"data_consulta"`
	Descricao    string `json:"descricao"`
}

type ConsultaResponseBody struct {
	ID           uint           `json:"id,omitempty"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
	PacienteID   uint           `json:"paciente_id,omitempty"`
	DentistaID   uint           `json:"dentista_id,omitempty"`
	DataConsulta time.Time      `json:"data_consulta,omitempty"`
	Descricao    string         `json:"descricao,omitempty"`
}
