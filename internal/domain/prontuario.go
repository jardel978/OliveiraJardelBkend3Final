package domain

import (
	"gorm.io/gorm"
	"time"
)

type PlanoTratamento struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"column:paciente_deleted_at;index"`
	ProntuarioID uint           `gorm:"foreignkey:ProntuarioID;not null"`
	Descricao    string         `gorm:"column:descricao;size=500"`
}

func (d *PlanoTratamento) TableName() string {
	return "plano_tratamentos"
}

type EvolucaoTratamento struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"column:paciente_deleted_at;index"`
	ProntuarioID uint           `gorm:"foreignkey:ProntuarioID;not null"`
	Descricao    string         `gorm:"column:descricao;size=500"`
}

func (d *EvolucaoTratamento) TableName() string {
	return "evolucao_tratamentos"
}

type Prontuario struct {
	gorm.Model
	PacienteID          uint `gorm:"foreignkey:PacienteID"`
	PlanosTratamento    []PlanoTratamento
	EvolucoesTratamento []EvolucaoTratamento
}

func (d *Prontuario) TableName() string {
	return "prontuarios"
}
