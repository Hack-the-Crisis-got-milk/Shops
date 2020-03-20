//+build wireinject

package main

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"github.com/Hack-the-Crisis-got-milk/Shops/routers"
	"github.com/Hack-the-Crisis-got-milk/Shops/utils"
	"github.com/google/wire"
)

func InitializeServer() (Server, error) {
	wire.Build(
		NewServer,
		routers.NewMainRouter,
		environment.NewEnv,
		utils.NewLogger,
	)
	return Server{}, nil
}
