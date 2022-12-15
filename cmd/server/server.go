package server

import (
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/routes"
	"OliveiraJardelBkend3Final/pkg/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

//Configurando o server

type Server struct {
	port   string
	server *gin.Engine
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Println("servidor rodando na porta:", s.port)
	log.Fatalln(router.Run(":" + s.port))
}

func NewServer() Server {
	severPort := configs.GetServerPort()
	s := Server{
		severPort,
		gin.Default(),
	}
	s.server.Use(middleware.Logger())
	return s
}
