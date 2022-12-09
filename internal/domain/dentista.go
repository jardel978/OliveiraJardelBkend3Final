package domain

import "gorm.io/gorm"

type Dentista struct {
	gorm.Model
	Nome         string     `gorm:"column:nome;size=20;not null"`
	Sobrenome    string     `gorm:"column:sobrenome;size=40;not null"`
	NomeCompleto string     `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
	EnderecoID   uint       `gorm:"column:endereco_id"`
	Matricula    string     `gorm:"column:matricula;size=20;not null"`
	Consultas    []Consulta `gorm:"column:consultas"`
}

func (d *Dentista) TableName() string {
	return "dentistas"
}
