package consulta

import (
	"OliveiraJardelBkend3Final/internal/dentista"
	"OliveiraJardelBkend3Final/internal/domain"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"OliveiraJardelBkend3Final/internal/paciente"
	"OliveiraJardelBkend3Final/internal/utils"
	"context"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
	"log"
	"time"
)

type CService interface {
	Save(consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (dtos.ConsultaResponseBody, error)
	SaveWithPacienteRgDentistaMatricula(pacienteRg string, dentistaMatricula string, consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (dtos.ConsultaResponseBody, error)
	FindAll(ctx context.Context) ([]dtos.ConsultaResponseBody, error)
	FindById(id uint, ctx context.Context) (dtos.ConsultaResponseBody, error)
	FindAllByRgPaciente(pacienteRg string, ctx context.Context) ([]dtos.ConsultaResponseBody, error)
	Update(id uint, consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (dtos.ConsultaResponseBody, error)
	Delete(id uint, ctx context.Context) error
}

type service struct {
	r  CRepository
	dr dentista.DRepository
	pr paciente.PRepository
}

func NewConsultaService() CService {
	return &service{
		r:  NewConsultaRepository(),
		dr: dentista.NewDentistaRepository(),
		pr: paciente.NewPacienteRepository(),
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

func (s *service) SaveWithPacienteRgDentistaMatricula(pacienteRg string, dentistaMatricula string, consultaDTO dtos.ConsultaRequestBody, ctx context.Context) (resp dtos.ConsultaResponseBody, err error) {
	var paciente domain.Paciente
	var dentista domain.Dentista

	paciente, err = s.pr.FindByRG(pacienteRg, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar paciente: paciente de rg:%v não encontrado.", pacienteRg),
			}
		}
		return resp, err
	}

	dentista, err = s.dr.FindByMatricula(dentistaMatricula, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar dentista: dentista de matrícula:%v não encontrado.", dentistaMatricula),
			}
		}
		return resp, err
	}

	consultaDTO.PacienteID = paciente.ID
	consultaDTO.DentistaID = dentista.ID

	return s.Save(consultaDTO, ctx)
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
				Message: fmt.Sprintf("falha ao buscar consulta: consulta de id:%v não encontrada.", id),
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

func (s *service) FindAllByRgPaciente(pacienteRg string, ctx context.Context) (resp []dtos.ConsultaResponseBody, err error) {
	var paciente domain.Paciente
	var list []domain.Consulta

	paciente, err = s.pr.FindByRG(pacienteRg, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar paciente: paciente de rg:%v não encontrado.", pacienteRg),
			}
		}
		return resp, err
	}

	list, err = s.r.FindAllByPacienteID(paciente.ID, ctx)
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
				Message: fmt.Sprintf("falha ao deletar consulta: consulta de id:%v não encontrada", id),
			}
		}
	}
	return err
}

func dtoToEntity(consultaDTO dtos.ConsultaRequestBody) (consulta domain.Consulta, err error) {

	mapper := dto.Mapper{}
	mapper.AddConvFunc(func(data string, mapper *dto.Mapper) time.Time {
		date, err := utils.ParseDataTime(data)
		if err != nil {
			log.Println(err)
		}
		return date
	})

	err = mapper.Map(&consulta, consultaDTO)
	if err != nil {
		return consulta, &errs.ErrInvalidMapping{Err: err}
	}

	return consulta, nil
}

func entityToDTO(consulta domain.Consulta) (resp dtos.ConsultaResponseBody, err error) {

	err = dto.Map(&resp, consulta)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}

	return resp, nil
}
