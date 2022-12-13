package dentista

import (
	"OliveiraJardelBkend3Final/internal/domain"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"context"
	"fmt"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
)

type DService interface {
	Save(dentistaDTO dtos.DentistaRequestBody, ctx context.Context) (resp dtos.DentistaResponseBody, err error)
	FindAll(ctx context.Context) (resp []dtos.DentistaResponseBody, err error)
	FindById(id uint, ctx context.Context) (resp dtos.DentistaResponseBody, err error)
	Update(id uint, dentistaDTO dtos.DentistaRequestBody, ctx context.Context) (resp dtos.DentistaResponseBody, err error)
	Delete(id uint, ctx context.Context) error
}

type service struct {
	r DRepository
}

func NewDentistaService() DService {
	return &service{
		r: NewDentistaRepository(),
	}
}

func (s *service) Save(dentistaDTO dtos.DentistaRequestBody, ctx context.Context) (resp dtos.DentistaResponseBody, err error) {
	dentista, errConvert := dtoToEntity(dentistaDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	dentista, err = s.r.Save(dentista, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(dentista)
	if errConvert != nil {
		return resp, errConvert
	}

	return resp, nil
}

func (s *service) FindAll(ctx context.Context) (resp []dtos.DentistaResponseBody, err error) {
	var list []domain.Dentista
	list, err = s.r.FindAll(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: "falha ao buscar dentistas: registros não encontrados.",
			}
		}
		return resp, err
	}

	if len(list) > 0 {
		for _, dentista := range list {
			dentistaDTO, errConvert := entityToDTO(dentista)
			if errConvert != nil {
				return resp, errConvert
			}
			resp = append(resp, dentistaDTO)
		}
	}

	return resp, nil
}

func (s *service) FindById(id uint, ctx context.Context) (resp dtos.DentistaResponseBody, err error) {
	var dentista domain.Dentista
	dentista, err = s.r.FindById(id, ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errs.ErrRecordNotFound{
				Message: fmt.Sprintf("falha ao buscar dentista: dentista de id:%v não encontrada.", id),
			}
		}
		return resp, err
	}

	resp, err = entityToDTO(dentista)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
func (s *service) Update(id uint, dentistaDTO dtos.DentistaRequestBody, ctx context.Context) (resp dtos.DentistaResponseBody, err error) {
	dentista, errConvert := dtoToEntity(dentistaDTO)
	if errConvert != nil {
		return resp, errConvert
	}

	dentista.DentistaID = id
	dentista, err = s.r.Update(dentista, ctx)
	if err != nil {
		return resp, err
	}

	resp, errConvert = entityToDTO(dentista)
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
				Message: fmt.Sprintf("falha ao deletar dentista: dentista de id:%v não encontrada", id),
			}
		}
	}
	return err
}

func dtoToEntity(dentistaDTO dtos.DentistaRequestBody) (dentista domain.Dentista, err error) {
	var endereco domain.Endereco

	err = dto.Map(&dentista, dentistaDTO)
	if err != nil {
		return dentista, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&endereco, dentistaDTO.EnderecoRequestBody)
	if err != nil {
		return dentista, &errs.ErrInvalidMapping{Err: err}
	}

	dentista.Endereco = endereco

	return dentista, nil
}

func entityToDTO(dentista domain.Dentista) (resp dtos.DentistaResponseBody, err error) {
	var enderecoResp dtos.EnderecoResponseBody

	err = dto.Map(&resp, dentista)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}
	err = dto.Map(&enderecoResp, dentista.Endereco)
	if err != nil {
		return resp, &errs.ErrInvalidMapping{Err: err}
	}

	resp.EnderecoResponseBody = enderecoResp

	return resp, nil
}
