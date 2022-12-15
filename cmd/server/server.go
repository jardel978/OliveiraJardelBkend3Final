package server

import (
	"OliveiraJardelBkend3Final/docs"
	"OliveiraJardelBkend3Final/internal/configs"
	"OliveiraJardelBkend3Final/internal/routes"
	"OliveiraJardelBkend3Final/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
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

func (s *Server) InitSwag() {
	s.server.Use(gin.Recovery(), gin.Logger())
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
	s.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
