package consulta

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
)

type CService interface {
	Save(consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (resp dtos.ConsultaResponseBody, err error)
	FindAll(ctx context.Context) (resp []dtos.ConsultaResponseBody, err error)
	FindById(id uint, ctx context.Context) (resp dtos.ConsultaResponseBody, err error)
	Update(id uint, consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (resp dtos.ConsultaResponseBody, err error)
	Delete(id uint, ctx context.Context) error
}

type service struct {
	r CRepository
}

func NewConsultaService() CService {
	return &service{
		r: NewConsultaRepository(),
	}
}

func (s *service) Save(consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (resp dtos.ConsultaResponseBody, err error) {
	consulta, errConvert := dtoToEntity(consultaDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	consulta, err = s.r.Save(consulta, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(consulta)
	if errConvert != nil {
		return resp, errConvert
	}

	return resp, nil
}

func (s *service) FindAll(ctx context.Context) (resp []dtos.ConsultaResponseBody, err error) {
	var list []domain.Consulta
	list, err = s.r.FindAll(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: "falha ao buscar consultas: registros não encontrados.",
			}
		}
		return resp, err
	}

	if len(list) > 0 {
		for _, consulta := range list {
			consultaDTO, errConvert := entityToDTO(consulta)
			if errConvert != nil {
				return resp, errConvert
			}
			resp = append(resp, consultaDTO)
		}
	}

	return resp, nil
}

func (s *service) FindById(id uint, ctx context.Context) (resp dtos.ConsultaResponseBody, err error) {
	var consulta domain.Consulta
	consulta, err = s.r.FindById(id, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar consulta: consulta de id:%v não encontrado.", id),
			}
		}
		return resp, err
	}

	resp, err = entityToDTO(consulta)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
func (s *service) Update(id uint, consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (resp dtos.ConsultaResponseBody, err error) {
	consulta, errConvert := dtoToEntity(consultaDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	consulta.ID = id
	consulta, err = s.r.Update(consulta, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(consulta)
	if errConvert != nil {
		return resp, errConvert
	}
	return resp, nil
}
func (s *service) Delete(id uint, ctx context.Context) error {
	err := s.r.Delete(id, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao deletar consulta: consulta de id:%v não encontrado", id),
			}
		}
	}
	return err
}

func dtoToEntity(consultaDTO dtos.ConsultaRequestBody) (consulta domain.Consulta, err error) {

	err = dto.Map(&consulta, consultaDTO)
	if err != nil {
		return consulta, &errs.ErrInvalidMapping{Err: err}
	}

	return consulta, nil
}

func entityToDTO(consulta domain.Consulta) (resp dtos.ConsultaResponseBody, err error) {
	var pacienteResp dtos.PacienteResponseBody
	var dentistaResp dtos.DentistaResponseBody
	err = dto.Map(&resp, consulta)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&pacienteResp, consulta.PacienteID)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&dentistaResp, consulta.DentistaID)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}

	resp.PacienteResponseBody = pacienteResp
	resp.DentistaResponseBody = dentistaResp

	return resp, nil
}
