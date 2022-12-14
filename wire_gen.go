// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"technical_test_skyshi/activity/controller"
	"technical_test_skyshi/activity/repository/mysql"
	"technical_test_skyshi/activity/service"
	"technical_test_skyshi/database"
	"technical_test_skyshi/router"
)

// Injectors from injector.go:

func InitServer() *gin.Engine {
	activityRepository := mysql.NewMysqlAcitivity()
	db := database.InitDBMysql()
	activityService := service.NewActivityServiceImpl(activityRepository, db)
	activityController := controller.NewActivityControllerImpl(activityService)
	engine := router.SetupRouter(activityController)
	return engine
}
