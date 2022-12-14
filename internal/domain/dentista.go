package domain

import "gorm.io/gorm"

type Dentista struct {
	gorm.Model
	Nome      string `gorm:"column:nome;size=20;not null"`
	Sobrenome string `gorm:"column:sobrenome;size=40;not null"`
	Endereco  `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Matricula string `gorm:"column:matricula;size=20;not null"`
	ClinicaID uint
	Consultas []Consulta
}

func (d *Dentista) TableName() string {
	return "dentistas"
}
