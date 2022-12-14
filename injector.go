//go:build wireinject
// +build wireinject

package main

import (
	"technical_test_skyshi/activity/controller"
	"technical_test_skyshi/activity/repository/mysql"
	"technical_test_skyshi/activity/service"
	"technical_test_skyshi/database"
	"technical_test_skyshi/router"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(
		database.InitDBMysql,

		mysql.NewMysqlAcitivity,
		service.NewActivityServiceImpl,
		controller.NewActivityControllerImpl,

		router.SetupRouter,
	)

	return nil
}
