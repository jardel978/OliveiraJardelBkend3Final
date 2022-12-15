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

// SalvarConsulta - Salva Consulta godoc
// @BasePath /api/v1
// @Summary Save Scheduling
// @Schemes
// @Tags Consultas
// @Description save scheduling in DB
// @Accept json
// @Produce json
// @Param consulta body handlers.consultaHandler true "Consulta Data"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consulta [post]
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
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 201, resp)
}

// SalvarConsultaComPacienteRgDentistaMatricula - Salva Consulta pelo RG do Paciente e Matricula do Dentista godoc
// @BasePath /api/v1
// @Summary Save Scheduling from enrollments
// @Tags Consultas
// @Description save scheduling from enrollments in DB
// @Accept json
// @Produce json
// @Param consulta body handlers.consultaHandler true "Consulta Data"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consulta/salvar [post]
func (ph *consultaHandler) SalvarConsultaComPacienteRgDentistaMatricula(ctx *gin.Context) {
	var dto dtos.ConsultaRequestBody
	errDtoJSON := ctx.ShouldBindJSON(&dto)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, &errs.ErrRequestBody{
			Err: errDtoJSON,
		})
		return
	}
	pacienteRg := ctx.Query("paciente-rg")
	dentistaMatricula := ctx.Query("dentista-matricula")

	resp, err := ph.s.SaveWithPacienteRgDentistaMatricula(pacienteRg, dentistaMatricula, dto, ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 201, resp)
}

// BuscarConsultas - Busca Todas as Consultas godoc
// @BasePath /api/v1
// @Summary Get All Schedulings
// @Tags Consultas
// @Description get all schedulings in DB
// @Accept json
// @Produce json
// @Param id path int64 true "Consultas ID"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consultas [get]
func (ph *consultaHandler) BuscarConsultas(ctx *gin.Context) {
	data, err := ph.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

// BuscarConsultaPorId - Busca Consulta por ID godoc
// @BasePath /api/v1
// @Summary Get Scheduling by ID
// @Tags Consultas
// @Description get scheduling by id in DB
// @Accept json
// @Produce json
// @Param id path int64 true "Consultas ID"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consultas/:id [get]
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

// BuscarConsultaPorPacienteRG - Busca Consulta por RG do paciente  godoc
// @BasePath /api/v1
// @Summary Get Scheduling by Patient Enrollment
// @Tags Consultas
// @Description get scheduling by patient enrollment in DB
// @Accept json
// @Produce json
// @Param rg path string true "Consultas RG"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consultas/pacientes/:rg [get]
func (ph *consultaHandler) BuscarConsultaPorPacienteRG(ctx *gin.Context) {
	pacienteRG := ctx.Param("rg")

	dto, err := ph.s.FindAllByRgPaciente(pacienteRG, ctx)
	if err != nil {
		web.Failure(ctx, 404, err)
		return
	}
	web.Success(ctx, 200, dto)
}

// AtualizarConsulta - Atualizar Consulta por ID godoc
// @BasePath /api/v1
// @Summary Put Scheduling by ID
// @Tags Consultas
// @Description get scheduling by ID in DB
// @Accept json
// @Produce json
// @Param consulta body handlers.consultaHandler true "Consulta Data"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consultas/:id [put]
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
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, resp)
}

// DeletarConsulta - Deletar Consulta por ID godoc
// @BasePath /api/v1
// @Summary Delete Scheduling by ID
// @Tags Consultas
// @Description delete scheduling by ID in DB
// @Accept json
// @Produce json
// @Param id path int64 true "Consultas ID"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /consultas/:id [delete]
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
