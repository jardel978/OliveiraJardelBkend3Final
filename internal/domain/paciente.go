package domain

import (
	"gorm.io/gorm"
	"time"
)

type Paciente struct {
	gorm.Model     `json:"gorm_._model"`
	Nome           string    `gorm:"column:nome;size=20;not null"`
	Sobrenome      string    `gorm:"column:sobrenome;size=40;not null"`
	NomeCompleto   string    `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
	DataNascimento time.Time `gorm:"column:data_nascimento;not null"`
	EnderecoID     uint      `gorm:"column:endereco_id"`
	ProntuarioID   uint      `gorm:"column:prontuario_id"`
	//DataCadastro   time.Time `json:"data_cadastro"`
}

func (d *Paciente) TableName() string {
	return "pacientes"
}
