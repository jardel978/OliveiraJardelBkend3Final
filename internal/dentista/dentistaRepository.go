package dentista

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type DRepository interface {
	Save(dentista domain.Dentista, ctx context.Context) (domain.Dentista, error)
	FindAll(ctx context.Context) ([]domain.Dentista, error)
	FindById(id uint, ctx context.Context) (domain.Dentista, error)
	FindByMatricula(rg string, ctx context.Context) (c domain.Dentista, err error)
	Update(dentista domain.Dentista, ctx context.Context) (domain.Dentista, error)
	Delete(id uint, ctx context.Context) error
}

type repository struct {
	repo IDentistaRepository
}

func NewDentistaRepository() DRepository {
	return &repository{repo: NewDentistaGorm()}
}

func (r *repository) Save(dentista domain.Dentista, ctx context.Context) (domain.Dentista, error) {
	return r.repo.Save(dentista, ctx)
}
func (r *repository) FindAll(ctx context.Context) ([]domain.Dentista, error) {
	return r.repo.FindAll(ctx)
}

func (r *repository) FindByMatricula(matricula string, ctx context.Context) (c domain.Dentista, err error) {
	return r.repo.FindByMatricula(matricula, ctx)
}

func (r *repository) FindById(id uint, ctx context.Context) (domain.Dentista, error) {
	return r.repo.FindById(id, ctx)
}
func (r *repository) Update(dentista domain.Dentista, ctx context.Context) (domain.Dentista, error) {
	return r.repo.Update(dentista, ctx)
}
func (r *repository) Delete(id uint, ctx context.Context) (err error) {
	return r.repo.Delete(id, ctx)
}
