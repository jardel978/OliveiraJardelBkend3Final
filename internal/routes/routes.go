package routes

import (
	"OliveiraJardelBkend3Final/cmd/server/handlers"
	"OliveiraJardelBkend3Final/pkg/web"
	"github.com/gin-gonic/gin"
)

//definindo rotas

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	clinicaHandler := handlers.NewClinicaHandler()

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
		//pacientes := main.Group("pacientes")
		//{
		//	pacientes.GET("/", handlers.BuscarPacientes)
		//	pacientes.GET("/:id", handlers.BuscarPacientePorId)
		//	pacientes.POST("/", handlers.SalvarPaciente)
		//	pacientes.PUT("/:id", handlers.AtualizarPaciente)
		//	pacientes.PATCH("/:id", handlers.AtualizarPaciente)
		//	pacientes.DELETE("/:id", handlers.DeletarPaciente)
		//}
	}
	return router
}
