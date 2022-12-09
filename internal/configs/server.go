package configs

import (
	routes2 "ddd-base/cmd/server/routes"
	"github.com/gin-gonic/gin"
	"log"
)

//Configurando o server

type Server struct {
	port   string
	server *gin.Engine
}

func (s *Server) Run() {
	router := routes2.ConfigRoutes(s.server)
	log.Println("server is runnig at port:", s.port)
	log.Fatalln(router.Run(":" + s.port))
}

func NewServer() Server {
	severPort := GetServerPort()
	return Server{
		severPort,
		gin.Default(),
	}
}
