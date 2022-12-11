package clinica

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type Repository interface {
	Save(clinica domain.Clinica, ctx context.Context) (domain.Clinica, error)
	FindAll(ctx context.Context) ([]domain.Clinica, error)
	FindById(id uint, ctx context.Context) (domain.Clinica, error)
	Update(clinica domain.Clinica, ctx context.Context) (domain.Clinica, error)
	Delete(id uint, ctx context.Context) error
}

type repository struct {
	repo IClinicaRepository
}

func NewClinicaRepository() Repository {
	return &repository{repo: NewClinicaGorm()}
}

func (r *repository) Save(clinica domain.Clinica, ctx context.Context) (domain.Clinica, error) {
	return r.repo.Save(clinica, ctx)
}
func (r *repository) FindAll(ctx context.Context) ([]domain.Clinica, error) {
	return r.repo.FindAll(ctx)
}

func (r *repository) FindById(id uint, ctx context.Context) (domain.Clinica, error) {
	return r.repo.FindById(id, ctx)
}
func (r *repository) Update(clinica domain.Clinica, ctx context.Context) (domain.Clinica, error) {
	return r.repo.Update(clinica, ctx)
}
func (r *repository) Delete(id uint, ctx context.Context) (err error) {
	return r.repo.Delete(id, ctx)
}
