package routes

import (
	"OliveiraJardelBkend3Final/cmd/server/handlers"
	"OliveiraJardelBkend3Final/pkg/web"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	clinicaHandler := handlers.NewClinicaHandler()
	pacienteHandler := handlers.NewPacienteHandler()
	dentistaHandler := handlers.NewDentistaHandler()
	consultaHandler := handlers.NewConsultaHandler()

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
		dentistas := main.Group("dentistas")
		{
			dentistas.GET("/", dentistaHandler.BuscarDentistas)
			dentistas.GET("/:id", dentistaHandler.BuscarDentistaPorId)
			dentistas.GET("/matriculas/:matricula", dentistaHandler.BuscarDentistaPorMatricula)
			dentistas.POST("/", dentistaHandler.SalvarDentista)
			dentistas.PUT("/:id", dentistaHandler.AtualizarDentista)
			dentistas.PATCH("/:id", dentistaHandler.AtualizarDentista)
			dentistas.DELETE("/:id", dentistaHandler.DeletarDentista)
		}

		consultas := main.Group("consultas")
		{
			consultas.GET("/", consultaHandler.BuscarConsultas)
			consultas.GET("/:id", consultaHandler.BuscarConsultaPorId)
			consultas.GET("/pacientes/:rg", consultaHandler.BuscarConsultaPorPacienteRG)
			consultas.POST("/", consultaHandler.SalvarConsulta)
			consultas.POST("/salvar", consultaHandler.SalvarConsultaComPacienteRgDentistaMatricula)
			consultas.PUT("/:id", consultaHandler.AtualizarConsulta)
			consultas.PATCH("/:id", consultaHandler.AtualizarConsulta)
			consultas.DELETE("/:id", consultaHandler.DeletarConsulta)
		}

	}
	return router
}
