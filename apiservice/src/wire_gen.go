// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minh1611/go_structure/apiservice/src/api"
	"github.com/minh1611/go_structure/apiservice/src/service"
)

// Injectors from server.wire.go:

func InitMainServer(ctx context.Context) (*Server, error) {
	engine := gin.New()
	playerService := service.PlayerService{}
	playerController := service.PlayerController{
		S: playerService,
	}
	apiService := &api.ApiService{
		G:      engine,
		Player: playerController,
	}
	server := &Server{
		ApiServer: apiService,
	}
	return server, nil
}

// server.wire.go:

type Server struct {
	ApiServer *api.ApiService
}