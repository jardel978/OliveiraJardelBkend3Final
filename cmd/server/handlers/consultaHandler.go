package handlers

import (
	"OliveiraJardelBkend3Final/internal/consulta"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"OliveiraJardelBkend3Final/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type consultaHandler struct {
	s consulta.CService
}

func NewConsultaHandler() *consultaHandler {
	return &consultaHandler{
		s: consulta.NewConsultaService(),
	}
}

func (ph *consultaHandler) SalvarConsulta(ctx *gin.Context) {
	var dto dtos.ConsultaRequestBody
	errDtoJSON := ctx.ShouldBindJSON(&dto)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, &errs.ErrRequestBody{
			Err: errDtoJSON,
		})
		return
	}
	resp, err := ph.s.Save(dto, ctx)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 201, resp)
}

func (ph *consultaHandler) BuscarConsultas(ctx *gin.Context) {
	data, err := ph.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

func (ph *consultaHandler) BuscarConsultaPorId(ctx *gin.Context) {
	id := ctx.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}
	dto, err := ph.s.FindById(uint(idNum), ctx)
	if err != nil {
		web.Failure(ctx, 404, err)
		return
	}
	web.Success(ctx, 200, dto)
}

func (ph *consultaHandler) AtualizarConsulta(ctx *gin.Context) {
	var dto dtos.ConsultaRequestBody
	id := ctx.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}

	errDtoJSON := ctx.ShouldBindJSON(&dto)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, &errs.ErrRequestBody{
			Err: errDtoJSON,
		})
		return
	}
	resp, err := ph.s.Update(uint(idNum), dto, ctx)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, resp)
}

func (ph *consultaHandler) DeletarConsulta(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}

	err = ph.s.Delete(uint(idNum), ctx)
	if err != nil {
		codeValue := 400
		if strings.Contains("não encontrada", err.Error()) {
			codeValue = 404
		}
		web.Failure(ctx, codeValue, err)
	}
}
