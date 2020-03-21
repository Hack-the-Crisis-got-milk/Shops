//+build wireinject

package main

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/config"
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"github.com/Hack-the-Crisis-got-milk/Shops/repositories"
	"github.com/Hack-the-Crisis-got-milk/Shops/routers"
	v1 "github.com/Hack-the-Crisis-got-milk/Shops/routers/api/v1"
	"github.com/Hack-the-Crisis-got-milk/Shops/utils"
	"github.com/google/wire"
)

func InitializeServer() (Server, error) {
	wire.Build(
		NewServer,
		routers.NewMainRouter,
		environment.NewEnv,
		utils.NewLogger,
		utils.NewDatabase,
		repositories.NewItemGroupRepository,
		v1.NewAPIV1Router,
		config.NewAppConfig,
	)
	return Server{}, nil
}
