package domain

import "gorm.io/gorm"

type PlanoTratamento struct {
	gorm.Model
	ProntuarioID uint   `json:"prontuario_id" gorm:"column:prontuario_id;not null"`
	Descricao    string `json:"descricao" gorm:"column:descricao;size=500"`
}

func (d *PlanoTratamento) TableName() string {
	return "plano_tratamentos"
}

type EvolucaoTratamento struct {
	gorm.Model
	ProntuarioID uint   `gorm:"column:prontuario_id;not null"`
	Descricao    string `gorm:"column:descricao;size=500"`
}

func (d *EvolucaoTratamento) TableName() string {
	return "evolucao_tratamentos"
}

type Prontuario struct {
	gorm.Model
	Consultas           []Consulta           `gorm:"column:consultas"`
	PlanosTratamento    []PlanoTratamento    `gorm:"column:planos_tratamento"`
	EvolucoesTratamento []EvolucaoTratamento `gorm:"column:evolucoes_tratamento"`
}

func (d *Prontuario) TableName() string {
	return "prontuarios"
}
