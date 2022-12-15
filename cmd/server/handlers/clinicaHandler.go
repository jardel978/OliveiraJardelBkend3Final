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

// SalvarClinica - Salva Clinica godoc
// @BasePath /api/v1
// @Summary Save Clinic
// @Schemes
// @Tags Clinica
// @Description save clinic in DB
// @Accept json
// @Produce json
// @Param clinica body handlers.clinicaHandler true "Clinicas Data"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /clinicas [post]
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

// BuscarClinicas - Buscar todas as Clinicas godoc
// @BasePath /api/v1
// @Summary List All Clinics
// @Tags Clinica
// @Description get all clinicas
// @Accept json
// @Produce json
// @Param id path int64 true "Clinicas ID"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /clinicas [get]
func (c *clinicaHandler) BuscarClinicas(ctx *gin.Context) {
	data, err := c.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

// BuscarClinicaPorId - Buscar Clinica por ID  godoc
// @BasePath /api/v1
// @Summary List Clinic by id
// @Tags Clinica
// @Description get clinics by id from db.
// @Accept json
// @Produce json
// @Param id path int64 true "Clinicas ID"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /clinicas/:id [get]
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

// AtualizarClinica - Atualizar Clinica por ID godoc
// @BasePath /api/v1
// @Summary Update Clinic
// @Tags Clinica
// @Description update clinic by ID
// @Accept json
// @Produce json
// @Param clinica body handlers.clinicaHandler true "Clinicas Data"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /clinica/:id [put]
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

// DeletarClinica - deletar Clinica por ID godoc
// @BasePath /api/v1
// @Summary Delete Clinic
// @Tags Clinica
// @Description delete clinic by ID
// @Accept json
// @Produce json
// @Param id path int64 true "Clinicas ID"
// @Success 200 {objet} map[string]interface{}
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /clinicas/:id [delete]
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
