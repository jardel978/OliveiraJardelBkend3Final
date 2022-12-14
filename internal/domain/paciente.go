package domain

import (
	"gorm.io/gorm"
	"time"
)

type Paciente struct {
	gorm.Model
	Nome           string    `gorm:"column:nome;size=20;not null"`
	Sobrenome      string    `gorm:"column:sobrenome;size=40;not null"`
	RG             string    `gorm:"column:rg;size=20;not null;unique"`
	DataNascimento time.Time `gorm:"column:data_nascimento;not null"`
	Prontuario     `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Endereco       `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Consultas      []Consulta
}

func (p *Paciente) TableName() string {
	return "pacientes"
}
