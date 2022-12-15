package handlers

import (
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"OliveiraJardelBkend3Final/internal/paciente"
	"OliveiraJardelBkend3Final/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type pacienteHandler struct {
	s paciente.PService
}

func NewPacienteHandler() *pacienteHandler {
	return &pacienteHandler{
		s: paciente.NewPacienteService(),
	}
}

// SalvarPaciente - Salva Paciente godoc
// @BasePath /api/v1
// @Summary Post Patient
// @Schemes
// @Tags Pacientes
// @Description post patients in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/ [post]
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
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 201, resp)
}

// BuscarPacientes - Busca todos  os Pacientes godoc
// @BasePath /api/v1
// @Summary Get All Patients
// @Tags Pacientes
// @Description get all patients in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes [get]
func (ph *pacienteHandler) BuscarPacientes(ctx *gin.Context) {
	data, err := ph.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

// BuscarPacientePorId - Busca Paciente por ID godoc
// @BasePath /api/v1
// @Summary Get Dentist by ID
// @Tags Pacientes
// @Description get dentist by ID in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [get]
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
		web.Failure(ctx, 404, err)
		return
	}
	web.Success(ctx, 200, dto)
}

// AtualizarPaciente - Atualiza Paciente por ID godoc
// @BasePath /api/v1
// @Summary Put Patient by ID
// @Tags Pacientes
// @Description put patients by ID in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes [put]
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
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, resp)
}

// DeletarPaciente - Deleta Paciente por ID godoc
// @BasePath /api/v1
// @Summary Deleta Patient by ID
// @Tags Pacientes
// @Description delete patients by ID in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [delete]
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
		if strings.Contains("não encontrado", err.Error()) {
			codeValue = 404
		}
		web.Failure(ctx, codeValue, err)
	}
}
