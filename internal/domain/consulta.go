package domain

import (
	"gorm.io/gorm"
	"time"
)

type Consulta struct {
	gorm.Model
	PacienteID uint      `gorm:"column:paciente_id;not null"`
	DentistaID uint      `gorm:"column:dentista_id;not null"`
	Data       time.Time `gorm:"column:data;not null"`
	Horario    time.Time `gorm:"column:horario;not null"`
	Descricao  string    `gorm:"column:descricao;size=250"`
}

func (d *Consulta) TableName() string {
	return "consultas"
}
