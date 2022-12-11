package handlers

//
//import (
//	"OliveiraJardelBkend3Final/internal/clinica"
//	"OliveiraJardelBkend3Final/internal/dtos"
//	"OliveiraJardelBkend3Final/internal/errs"
//	"OliveiraJardelBkend3Final/pkg/web"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"strconv"
//	"strings"
//)
//
//var sPaciente clinica.Service
//
//func init() {
//	sPaciente = clinica.NewClinicaService()
//}
//
//func SalvarPaciente(ctx *gin.Context) {
//	var dto dtos.ClinicaRequestBody
//	errDtoJSON := ctx.ShouldBindJSON(&dto)
//	if errDtoJSON != nil {
//		web.Failure(ctx, 400, &errs.ErrRequestBody{
//			Err: errDtoJSON,
//		})
//		return
//	}
//	resp, err := sClinic.Save(dto, ctx)
//	if errDtoJSON != nil {
//		web.Failure(ctx, 400, err)
//		return
//	}
//	web.Success(ctx, 201, resp)
//}
//
//func BuscarPacientes(ctx *gin.Context) {
//	data, err := sClinic.FindAll(ctx)
//	if err != nil {
//		web.Failure(ctx, 400, err)
//		return
//	}
//	web.Success(ctx, 200, data)
//}
//
//func BuscarPacientePorId(ctx *gin.Context) {
//	id := ctx.Param("id")
//
//	idNum, err := strconv.Atoi(id)
//	if err != nil {
//		web.Failure(ctx, 400, &errs.ErrInvalidParams{
//			Err: err,
//		})
//		return
//	}
//	dto, err := sClinic.FindById(uint(idNum), ctx)
//	if err != nil {
//		web.Failure(ctx, 404, &errs.ErrRecordNotFound{Message: fmt.Sprintf("clínica de id: %v não encontrada.", idNum)})
//		return
//	}
//	web.Success(ctx, 200, dto)
//}
//
//func AtualizarPaciente(ctx *gin.Context) {
//	var dto dtos.ClinicaRequestBody
//	id := ctx.Param("id")
//	idNum, err := strconv.Atoi(id)
//	if err != nil {
//		web.Failure(ctx, 400, &errs.ErrInvalidParams{
//			Err: err,
//		})
//		return
//	}
//
//	errDtoJSON := ctx.ShouldBindJSON(&dto)
//	if errDtoJSON != nil {
//		web.Failure(ctx, 400, &errs.ErrRequestBody{
//			Err: errDtoJSON,
//		})
//		return
//	}
//	resp, err := sClinic.Update(uint(idNum), dto, ctx)
//	if errDtoJSON != nil {
//		web.Failure(ctx, 400, err)
//		return
//	}
//	web.Success(ctx, 200, resp)
//}
//
//func DeletarPaciente(ctx *gin.Context) {
//	id := ctx.Param("id")
//	idNum, err := strconv.Atoi(id)
//	if err != nil {
//		web.Failure(ctx, 400, &errs.ErrInvalidParams{
//			Err: err,
//		})
//		return
//	}
//
//	err = sClinic.Delete(uint(idNum), ctx)
//	if err != nil {
//		codeValue := 400
//		if strings.Contains("não encontrada", err.Error()) {
//			codeValue = 404
//		}
//		web.Failure(ctx, codeValue, err)
//	}
//}
