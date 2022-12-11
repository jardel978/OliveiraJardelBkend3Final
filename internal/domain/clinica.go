package domain

import "gorm.io/gorm"

type Clinica struct {
	gorm.Model
	Nome      string `gorm:"column:nome;size:50;not null"`
	CNPJ      string `gorm:"column:cnpj;size:18;unique;not null"`
	Endereco  Endereco
	Dentistas []Dentista `gorm:"many2many:clinica_dentistas"`
}

func (c *Clinica) TableName() string {
	return "clinicas"
}
