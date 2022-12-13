package paciente

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type IPacienteRepository interface {
	Save(paciente domain.Paciente, ctx context.Context) (c domain.Paciente, err error)
	FindAll(ctx context.Context) (list []domain.Paciente, err error)
	FindById(id uint, ctx context.Context) (c domain.Paciente, err error)
	FindByRG(rg string, ctx context.Context) (c domain.Paciente, err error)
	Update(paciente domain.Paciente, ctx context.Context) (c domain.Paciente, err error)
	Delete(id uint, ctx context.Context) (err error)
}
