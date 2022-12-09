package domain

import "gorm.io/gorm"

type Clinica struct {
	gorm.Model
	Nome       string     `gorm:"column:nome;size:50;not null"`
	CNPJ       string     `gorm:"column:cnpj;size:18;unique;not null"`
	EnderecoID uint       `gorm:"column:endereco_id"`
	Endereco   Endereco   `gorm:"column:endereco"`
	Dentistas  []Dentista `gorm:"column:dentistas"`
}

func (d *Clinica) TableName() string {
	return "clinicas"
}
