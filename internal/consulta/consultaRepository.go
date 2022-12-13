package consulta

import "OliveiraJardelBkend3Final/internal/domain"

type CRepository interface {
	Save(consulta domain.Consulta, ctx context.Context) (domain.Consulta, error)
	FindAll(ctx context.Context) ([]domain.Consulta, error)
	FindById(id uint, ctx context.Context) (domain.Consulta, error)
	FindByRG(rg string, ctx context.Context) (c domain.Consulta, err error)
	Update(consulta domain.Consulta, ctx context.Context) (domain.Consulta, error)
	Delete(id uint, ctx context.Context) error
}

type repository struct {
	repo IConsultaRepository
}

func NewConsultaRepository() CRepository {
	return &repository{repo: NewConsultaGorm()}
}

func (r *repository) Save(consulta domain.Consulta, ctx context.Context) (domain.Consulta, error) {
	return r.repo.Save(consulta, ctx)
}
func (r *repository) FindAll(ctx context.Context) ([]domain.Consulta, error) {
	return r.repo.FindAll(ctx)
}

func (r *repository) FindById(id uint, ctx context.Context) (domain.Consulta, error) {
	return r.repo.FindById(id, ctx)
}
func (r *repository) Update(consulta domain.Consulta, ctx context.Context) (domain.Consulta, error) {
	return r.repo.Update(consulta, ctx)
}
func (r *repository) Delete(id uint, ctx context.Context) (err error) {
	return r.repo.Delete(id, ctx)
}
