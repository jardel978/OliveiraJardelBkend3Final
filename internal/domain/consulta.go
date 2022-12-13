package domain

import (
	"gorm.io/gorm"
	"time"
)

type Consulta struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"column:paciente_deleted_at;index"`
	PacienteID uint           `gorm:"not null"`
	//Paciente
	DentistaID uint `gorm:"not null"`
	//Dentista
	Data      time.Time `gorm:"column:data;not null"`
	Horario   time.Time `gorm:"column:horario;not null"`
	Descricao string    `gorm:"column:descricao;size=250"`
}

func (d *Consulta) TableName() string {
	return "consultas"
}
