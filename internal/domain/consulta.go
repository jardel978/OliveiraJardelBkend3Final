package domain

import (
	"gorm.io/gorm"
	"time"
)

type Consulta struct {
	gorm.Model
	PacienteID   uint      `gorm:"not null"`
	DentistaID   uint      `gorm:"not null"`
	DataConsulta time.Time `gorm:"column:data_consulta;not null"`
	Descricao    string    `gorm:"column:descricao;size=250"`
}

func (d *Consulta) TableName() string {
	return "consultas"
}
