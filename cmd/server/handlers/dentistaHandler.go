package handlers

import (
	"OliveiraJardelBkend3Final/internal/dentista"
	"OliveiraJardelBkend3Final/internal/dtos"
	"OliveiraJardelBkend3Final/internal/errs"
	"OliveiraJardelBkend3Final/pkg/web"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

type dentistaHandler struct {
	s dentista.DService
}

func NewDentistaHandler() *dentistaHandler {
	return &dentistaHandler{
		s: dentista.NewDentistaService(),
	}
}

// SalvarDentista - Salva Dentista godoc
// @BasePath /api/v1
// @Summary Save Dentist
// @Tags Dentistas
// @Description save dentist in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas [post]
func (dh *dentistaHandler) SalvarDentista(ctx *gin.Context) {
	var dto dtos.DentistaRequestBody
	errDtoJSON := ctx.ShouldBindJSON(&dto)
	if errDtoJSON != nil {
		web.Failure(ctx, 400, &errs.ErrRequestBody{
			Err: errDtoJSON,
		})
		return
	}
	log.Printf("\nhandler DTO: %v", dto)
	resp, err := dh.s.Save(dto, ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 201, resp)
}

// BuscarDentistas - Busca Dentistas godoc
// @BasePath /api/v1
// @Summary Get All Dentists
// @Tags Dentistas
// @Description get all dentists in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas [get]
func (dh *dentistaHandler) BuscarDentistas(ctx *gin.Context) {
	data, err := dh.s.FindAll(ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, data)
}

// BuscarDentistaPorId - Busca um Dentista por ID godoc
// @BasePath /api/v1
// @Summary get Dentist by ID
// @Tags Dentistas
// @Description get dentist by id in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas/:i [get]
func (dh *dentistaHandler) BuscarDentistaPorId(ctx *gin.Context) {
	id := ctx.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}
	dto, err := dh.s.FindById(uint(idNum), ctx)
	if err != nil {
		web.Failure(ctx, 404, err)
		return
	}
	web.Success(ctx, 200, dto)
}

// BuscarDentistaPorMatricula - Busca Dentista por Matricula godoc
// @BasePath /api/v1
// @Summary Get Dentist by Enrollment
// @Tags Dentistas
// @Description get dentist by enrollment in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas/matriculas/:matricula [get]
func (dh *dentistaHandler) BuscarDentistaPorMatricula(ctx *gin.Context) {
	matricula := ctx.Param("matricula")

	dto, err := dh.s.FindByMatricula(matricula, ctx)
	if err != nil {
		web.Failure(ctx, 404, err)
		return
	}
	web.Success(ctx, 200, dto)
}

// AtualizarDentista - Atualiza Dentista godoc
// @BasePath /api/v1
// @Summary Put Dentist by ID
// @Schemes
// @Tags Dentistas
// @Description put dentist by ID in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas/:id [put]
func (dh *dentistaHandler) AtualizarDentista(ctx *gin.Context) {
	var dto dtos.DentistaRequestBody
	id := ctx.Param("id")
	idNum, errId := strconv.Atoi(id)
	if errId != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: errId,
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
	resp, err := dh.s.Update(uint(idNum), dto, ctx)
	if err != nil {
		web.Failure(ctx, 400, err)
		return
	}
	web.Success(ctx, 200, resp)
}

// DeletarDentista - Deleta Dentista godoc
// @BasePath /api/v1
// @Summary Deleta Dentist
// @Tags Dentistas
// @Description deleta dentist in DB
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {objet} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas/:id [delete]
func (dh *dentistaHandler) DeletarDentista(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		web.Failure(ctx, 400, &errs.ErrInvalidParams{
			Err: err,
		})
		return
	}

	err = dh.s.Delete(uint(idNum), ctx)
	if err != nil {
		codeValue := 400
		if strings.Contains("não encontrado", err.Error()) {
			codeValue = 404
		}
		web.Failure(ctx, codeValue, err)
	}
}
