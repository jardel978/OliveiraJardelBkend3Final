package domain

import (
	"gorm.io/gorm"
	"time"
)

type Dentista struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"column:paciente_deleted_at;index"`
	Nome      string         `gorm:"column:nome;size=20;not null"`
	Sobrenome string         `gorm:"column:sobrenome;size=40;not null"`
	Endereco
	Matricula string `gorm:"column:matricula;size=20;not null"`
	Consultas []Consulta
	Clinicas  []Clinica `gorm:"many2many:dentistas_clinicas"`
}

func (d *Dentista) TableName() string {
	return "dentistas"
}
