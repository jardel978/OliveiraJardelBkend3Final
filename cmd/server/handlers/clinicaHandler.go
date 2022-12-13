package handlers

import (
	"OliveiraJardelBkend3Final/internal/clinica"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"OliveiraJardelBkend3Final/pkg/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type clinicaHandler struct {
	s clinica.CService
}

func NewClinicaHandler() *clinicaHandler {
	return &clinicaHandler{
		s: clinica.NewClinicaService(),
	}
}

func (c *clinicaHandler) SalvarClinica(ctx *gin.Context) {
	var dto dtos.ClinicaRequestBody
	errDtoJSON := ctx.ShouldBindJSON(&dto)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, &errs.ErrRequestBody{
			Err: errDtoJSON,
		})
		return
	}
	resp, err := c.s.Save(dto, ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 201, resp)
}

func (c *clinicaHandler) BuscarClinicas(ctx *gin.Context) {
	data, err := c.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

func (c *clinicaHandler) BuscarClinicaPorId(ctx *gin.Context) {
	id := ctx.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}
	dto, err := c.s.FindById(uint(idNum), ctx)
	if err != nil {
		web.Failure(ctx, 404, &errs.ErrRecordNotFound{Message: fmt.Sprintf("clínica de id: %v não encontrada.", idNum)})
		return
	}
	web.Success(ctx, 200, dto)
}

func (c *clinicaHandler) AtualizarClinica(ctx *gin.Context) {
	var dto dtos.ClinicaRequestBody
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
	resp, err := c.s.Update(uint(idNum), dto, ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, resp)
}

func (c *clinicaHandler) DeletarClinica(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}

	err = c.s.Delete(uint(idNum), ctx)
	if err != nil {
		codeValue := 400
		if strings.Contains("não encontrada", err.Error()) {
			codeValue = 404
		}
		web.Failure(ctx, codeValue, err)
	}
}
