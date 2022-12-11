package domain

import "gorm.io/gorm"

type PlanoTratamento struct {
	gorm.Model
	ProntuarioID uint   `gorm:"not null"`
	Descricao    string `gorm:"column:descricao;size=500"`
}

func (d *PlanoTratamento) TableName() string {
	return "plano_tratamentos"
}

type EvolucaoTratamento struct {
	gorm.Model
	ProntuarioID uint   `gorm:"not null"`
	Descricao    string `gorm:"column:descricao;size=500"`
}

func (d *EvolucaoTratamento) TableName() string {
	return "evolucao_tratamentos"
}

type Prontuario struct {
	gorm.Model
	PacienteID          uint
	PlanosTratamento    []PlanoTratamento
	EvolucoesTratamento []EvolucaoTratamento
}

func (d *Prontuario) TableName() string {
	return "prontuarios"
}
