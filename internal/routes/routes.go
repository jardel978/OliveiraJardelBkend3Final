package routes

import (
	"OliveiraJardelBkend3Final/cmd/server/handlers"
	"OliveiraJardelBkend3Final/pkg/web"
	"github.com/gin-gonic/gin"
)

//definindo rotas

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	clinicaHandler := handlers.NewClinicaHandler()
	pacienteHandler := handlers.NewPacienteHandler()

	main := router.Group("api/v1") // base url
	main.GET("/ping", func(ctx *gin.Context) {
		web.Success(ctx, 200, "pong")
	})
	{
		clinicas := main.Group("clinicas")
		{
			clinicas.GET("/", clinicaHandler.BuscarClinicas)
			clinicas.GET("/:id", clinicaHandler.BuscarClinicaPorId)
			clinicas.POST("/", clinicaHandler.SalvarClinica)
			clinicas.PUT("/:id", clinicaHandler.AtualizarClinica)
			clinicas.PATCH("/:id", clinicaHandler.AtualizarClinica)
			clinicas.DELETE("/:id", clinicaHandler.DeletarClinica)
		}
		pacientes := main.Group("pacientes")
		{
			pacientes.GET("/", pacienteHandler.BuscarPacientes)
			pacientes.GET("/:id", pacienteHandler.BuscarPacientePorId)
			pacientes.POST("/", pacienteHandler.SalvarPaciente)
			pacientes.PUT("/:id", pacienteHandler.AtualizarPaciente)
			pacientes.PATCH("/:id", pacienteHandler.AtualizarPaciente)
			pacientes.DELETE("/:id", pacienteHandler.DeletarPaciente)
		}
	}
	return router
}
