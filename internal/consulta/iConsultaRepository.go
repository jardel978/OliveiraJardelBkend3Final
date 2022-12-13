package consulta

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type IConsultaRepository interface {
	Save(consulta domain.Consulta, ctx context.Context) (domain.Consulta, error)
	FindAll(ctx context.Context) ([]domain.Consulta, error)
	FindById(id uint, ctx context.Context) (domain.Consulta, error)
	FindAllByPacienteID(pacienteID uint, ctx context.Context) ([]domain.Consulta, error)
	Update(consulta domain.Consulta, ctx context.Context) (domain.Consulta, error)
	Delete(id uint, ctx context.Context) error
}
