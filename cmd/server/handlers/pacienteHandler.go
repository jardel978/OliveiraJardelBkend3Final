package handlers

import (
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"OliveiraJardelBkend3Final/internal/paciente"
	"OliveiraJardelBkend3Final/pkg/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type pacienteHandler struct {
	s paciente.PService
}

func NewPacienteHandler() *pacienteHandler {
	return &pacienteHandler{
		s: paciente.NewPacinteService(),
	}
}

func (ph *pacienteHandler) SalvarPaciente(ctx *gin.Context) {
	var dto dtos.PacienteRequestBody
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

func (ph *pacienteHandler) BuscarPacientes(ctx *gin.Context) {
	data, err := ph.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

func (ph *pacienteHandler) BuscarPacientePorId(ctx *gin.Context) {
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
		web.Failure(ctx, 404, &errs.ErrRecordNotFound{Message: fmt.Sprintf("clínica de id: %v não encontrada.", idNum)})
		return
	}
	web.Success(ctx, 200, dto)
}

func (ph *pacienteHandler) AtualizarPaciente(ctx *gin.Context) {
	var dto dtos.PacienteRequestBody
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

func (ph *pacienteHandler) DeletarPaciente(ctx *gin.Context) {
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
