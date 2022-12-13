package consulta

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type IConsultaRepository interface {
	Save(consulta domain.Consulta, ctx context.Context) (c domain.Consulta, err error)
	FindAll(ctx context.Context) (list []domain.Consulta, err error)
	FindById(id uint, ctx context.Context) (c domain.Consulta, err error)
	Update(consulta domain.Consulta, ctx context.Context) (c domain.Consulta, err error)
	Delete(id uint, ctx context.Context) (err error)
}
