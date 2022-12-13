package domain

import (
	"gorm.io/gorm"
	"time"
)

type Paciente struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"column:paciente_deleted_at;index"`
	Nome           string         `gorm:"column:nome;size=20;not null"`
	Sobrenome      string         `gorm:"column:sobrenome;size=40;not null"`
	RG             string         `gorm:"column:rg;size=20;not null;unique"`
	DataNascimento time.Time      `gorm:"column:data_nascimento;not null"`
	Consultas      []Consulta
	Prontuario     `gorm:"foreignkey:ProntuarioID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	Endereco       `gorm:"foreignkey:EnderecoID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}

func (p *Paciente) TableName() string {
	return "pacientes"
}
