package clinica

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type IClinicaRepository interface {
	Save(clinica domain.Clinica, ctx context.Context) (c domain.Clinica, err error)
	FindAll(ctx context.Context) (list []domain.Clinica, err error)
	FindById(id uint, ctx context.Context) (c domain.Clinica, err error)
	Update(clinica domain.Clinica, ctx context.Context) (c domain.Clinica, err error)
	Delete(id uint, ctx context.Context) (err error)
}
