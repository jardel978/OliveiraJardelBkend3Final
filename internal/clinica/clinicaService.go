package clinica

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"context"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
	"log"
)

type CService interface {
	Save(clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error)
	FindAll(ctx context.Context) (resp []dtos.ClinicaResponseBody, err error)
	FindById(id uint, ctx context.Context) (resp dtos.ClinicaResponseBody, err error)
	Update(id uint, clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error)
	Delete(id uint, ctx context.Context) error
}

type service struct {
	r CRepository
}

func NewClinicaService() CService {
	return &service{
		r: NewClinicaRepository(),
	}
}

func (s *service) Save(clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error) {
	clinica, errConvert := dtoToEntity(clinicaDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	clinica, err = s.r.Save(clinica, ctx)
	if err != nil {
		return resp, err
	}
	//clinica.EnderecoID = clinica.Endereco.ID
	//clinica, err = s.r.Update(clinica, ctx)
	//if err != nil {
	//	return resp, err
	//}

	resp, errConvert = entityToDTO(clinica)
	if errConvert != nil {
		return resp, errConvert
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

	log.Printf("\nlist - (%d) - clinicas: %v\n", len(list), list)

	if len(list) > 0 {
		for _, clinica := range list {
			clinicaDTO, errConvert := entityToDTO(clinica)
			if errConvert != nil {
				return resp, errConvert
			}
			resp = append(resp, clinicaDTO)
		}
	}

	log.Printf("\nresp clinicas: %v\n", resp)
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
	fmt.Printf("clinica get: %v\n", clinica)

	resp, err = entityToDTO(clinica)
	if err != nil {
		return resp, err
	}
	fmt.Printf("clinicadto: %v\n", resp)

	return resp, nil
}
func (s *service) Update(id uint, clinicaDTO dtos.ClinicaRequestBody, ctx context.Context) (resp dtos.ClinicaResponseBody, err error) {
	clinica, errConvert := dtoToEntity(clinicaDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	clinica.ID = id

	clinica, err = s.r.Update(clinica, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(clinica)
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
				Message: fmt.Sprintf("falha ao deletar clinica: clinica de id:%v não encontrada", id),
			}
		}
	}
	return err
}

func dtoToEntity(clinicaDTO dtos.ClinicaRequestBody) (clinica domain.Clinica, err error) {
	var endereco domain.Endereco

	err = dto.Map(&clinica, clinicaDTO)
	if err != nil {
		return clinica, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&endereco, clinicaDTO.EnderecoRequestBody)
	if err != nil {
		return clinica, &errs.ErrInvalidMapping{Err: err}
	}
	clinica.Endereco = endereco

	return clinica, nil
}

func entityToDTO(clinica domain.Clinica) (resp dtos.ClinicaResponseBody, err error) {
	var enderecoResp dtos.EnderecoResponseBody
	err = dto.Map(&resp, clinica)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&enderecoResp, clinica.Endereco)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}

	resp.EnderecoResponseBody = enderecoResp

	return resp, nil
}
