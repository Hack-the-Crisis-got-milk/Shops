package main

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"github.com/Hack-the-Crisis-got-milk/Shops/routers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	Port string
}

func NewServer(mainRouter routers.MainRouter, env *environment.Env) Server {
	server := Server{
		Engine: gin.Default(),
		Port:   env.Get(environment.Port),
	}

	mainRouter.RegisterRoutes(server.Group("/"))

	return server
}
