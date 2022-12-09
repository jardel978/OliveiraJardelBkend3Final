package dtos

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"gorm.io/gorm"
	"time"
)

type ConsultaRequestBody struct {
	PacienteID uint      `json:"paciente_id" binding:"required"`
	DentistaID uint      `json:"dentista_id" binding:"required"`
	Data       time.Time `json:"data" binding:"required"`
	Horario    time.Time `json:"horario" binding:"required"`
	Descricao  string    `json:"descricao"`
}

type ConsultaResponseBody struct {
	ID              uint           `json:"id"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
	domain.Paciente `json:"paciente"`
	domain.Dentista `json:"dentista"`
	Data            time.Time `json:"data"`
	Horario         time.Time `json:"horario"`
	Descricao       string    `json:"descricao,omitempty"`
}
