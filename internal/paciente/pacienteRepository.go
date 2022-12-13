package paciente

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"context"
)

type PRepository interface {
	Save(paciente domain.Paciente, ctx context.Context) (domain.Paciente, error)
	FindAll(ctx context.Context) ([]domain.Paciente, error)
	FindById(id uint, ctx context.Context) (domain.Paciente, error)
	FindByRG(rg string, ctx context.Context) (c domain.Paciente, err error)
	Update(paciente domain.Paciente, ctx context.Context) (domain.Paciente, error)
	Delete(id uint, ctx context.Context) error
}

type repository struct {
	repo IPacienteRepository
}

func NewPacienteRepository() PRepository {
	return &repository{repo: NewPacienteGorm()}
}

func (r *repository) Save(paciente domain.Paciente, ctx context.Context) (domain.Paciente, error) {
	return r.repo.Save(paciente, ctx)
}
func (r *repository) FindAll(ctx context.Context) ([]domain.Paciente, error) {
	return r.repo.FindAll(ctx)
}

func (r *repository) FindByRG(rg string, ctx context.Context) (c domain.Paciente, err error) {
	return r.repo.FindByRG(rg, ctx)
}

func (r *repository) FindById(id uint, ctx context.Context) (domain.Paciente, error) {
	return r.repo.FindById(id, ctx)
}
func (r *repository) Update(paciente domain.Paciente, ctx context.Context) (domain.Paciente, error) {
	return r.repo.Update(paciente, ctx)
}
func (r *repository) Delete(id uint, ctx context.Context) (err error) {
	return r.repo.Delete(id, ctx)
}
