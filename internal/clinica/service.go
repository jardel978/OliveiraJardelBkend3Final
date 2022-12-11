package clinica

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"context"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
)

type Service interface {
	Save(clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error)
	FindAll(ctx context.Context) (resp []dtos.ClinicaResponseBody, err error)
	FindById(id uint, ctx context.Context) (resp dtos.ClinicaResponseBody, err error)
	Update(id uint, clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error)
	Delete(id uint, ctx context.Context) error
}

type service struct {
	r Repository
}

func NewClinicaService() Service {
	return &service{
		r: NewClinicaRepository(),
	}
}

func (s *service) Save(clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error) {
	var clinica domain.Clinica
	err = dto.Map(&clinica, clinicaDTO)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	clinica, err = s.r.Save(clinica, ctx)
	if err != nil {
		return resp, err
	}
	err = dto.Map(&resp, clinica)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}

	return resp, nil
}

func (s *service) FindAll(ctx context.Context) (resp []dtos.ClinicaResponseBody, err error) {
	var list []domain.Clinica
	list, err = s.r.FindAll(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: "falha ao buscar clinicas: registros não encontrados.",
			}
		}
		return resp, err
	}
	err = dto.Map(&resp, list)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	return resp, nil
}

func (s *service) FindById(id uint, ctx context.Context) (resp dtos.ClinicaResponseBody, err error) {
	var clinica domain.Clinica
	clinica, err = s.r.FindById(id, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar clinica: clinica de id:%v não encontrada.", id),
			}
		}
		return resp, err
	}
	err = dto.Map(&resp, clinica)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	return resp, nil
}
func (s *service) Update(id uint, clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error) {
	var clinica domain.Clinica
	err = dto.Map(&clinica, clinicaDTO)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	clinica.ID = id
	clinica, err = s.r.Update(clinica, ctx)
	if err != nil {
		return resp, err
	}
	err = dto.Map(&resp, clinica)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	return resp, nil
}
func (s *service) Delete(id uint, ctx context.Context) error {
	err := s.r.Delete(id, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao deletar clinica: clinica de id:%v não encontrada", id),
			}
		}
	}
	return err
}
