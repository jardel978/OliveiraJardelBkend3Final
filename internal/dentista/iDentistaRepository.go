package dentista

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type IDentistaRepository interface {
	Save(dentista domain.Dentista, ctx context.Context) (c domain.Dentista, err error)
	FindAll(ctx context.Context) (list []domain.Dentista, err error)
	FindById(id uint, ctx context.Context) (c domain.Dentista, err error)
	FindByMatricula(matricula string, ctx context.Context) (c domain.Dentista, err error)
	Update(dentista domain.Dentista, ctx context.Context) (c domain.Dentista, err error)
	Delete(id uint, ctx context.Context) (err error)
}
