package paciente

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"context"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
)

type PService interface {
	Save(pacienteDTO dtos.PacienteRequestBody, ctx context.Context) (resp dtos.PacienteResponseBody, err error)
	FindAll(ctx context.Context) (resp []dtos.PacienteResponseBody, err error)
	FindById(id uint, ctx context.Context) (resp dtos.PacienteResponseBody, err error)
	Update(id uint, pacienteDTO dtos.PacienteRequestBody, ctx context.Context) (resp dtos.PacienteResponseBody, err error)
	Delete(id uint, ctx context.Context) error
}

type service struct {
	r PRepository
}

func NewPacinteService() PService {
	return &service{
		r: NewPacienteRepository(),
	}
}

func (s *service) Save(pacienteDTO dtos.PacienteRequestBody, ctx context.Context) (resp dtos.PacienteResponseBody, err error) {
	paciente, errConvert := dtoToEntity(pacienteDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	paciente, err = s.r.Save(paciente, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(paciente)
	if errConvert != nil {
		return resp, errConvert
	}

	return resp, nil
}

func (s *service) FindAll(ctx context.Context) (resp []dtos.PacienteResponseBody, err error) {
	var list []domain.Paciente
	list, err = s.r.FindAll(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: "falha ao buscar pacientes: registros não encontrados.",
			}
		}
		return resp, err
	}

	if len(list) > 0 {
		for _, paciente := range list {
			pacienteDTO, errConvert := entityToDTO(paciente)
			if errConvert != nil {
				return resp, errConvert
			}
			resp = append(resp, pacienteDTO)
		}
	}

	return resp, nil
}

func (s *service) FindById(id uint, ctx context.Context) (resp dtos.PacienteResponseBody, err error) {
	var paciente domain.Paciente
	paciente, err = s.r.FindById(id, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar paciente: paciente de id:%v não encontrado.", id),
			}
		}
		return resp, err
	}

	resp, err = entityToDTO(paciente)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
func (s *service) Update(id uint, pacienteDTO dtos.PacienteRequestBody, ctx context.Context) (resp dtos.PacienteResponseBody, err error) {
	paciente, errConvert := dtoToEntity(pacienteDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	paciente.ID = id
	paciente, err = s.r.Update(paciente, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(paciente)
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
				Message: fmt.Sprintf("falha ao deletar paciente: paciente de id:%v não encontrado", id),
			}
		}
	}
	return err
}

func dtoToEntity(pacienteDTO dtos.PacienteRequestBody) (paciente domain.Paciente, err error) {
	var endereco domain.Endereco
	var prontuario domain.Prontuario

	err = dto.Map(&paciente, pacienteDTO)
	if err != nil {
		return paciente, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&endereco, pacienteDTO.EnderecoRequestBody)
	if err != nil {
		return paciente, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&prontuario, pacienteDTO.ProntuarioRequestBody)
	if err != nil {
		return paciente, &errs.ErrInvalidMapping{Err: err}
	}
	paciente.Endereco = endereco
	paciente.Prontuario = prontuario

	return paciente, nil
}

func entityToDTO(paciente domain.Paciente) (resp dtos.PacienteResponseBody, err error) {
	var enderecoResp dtos.EnderecoResponseBody
	var prontuarioResp dtos.ProntuarioResponseBody
	err = dto.Map(&resp, paciente)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&enderecoResp, paciente.Endereco)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&prontuarioResp, paciente.Prontuario)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}

	resp.EnderecoResponseBody = enderecoResp
	resp.ProntuarioResponseBody = prontuarioResp

	return resp, nil
}
