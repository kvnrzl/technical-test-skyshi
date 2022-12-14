//go:build wireinject
// +build wireinject

package main

import (
	activityController "technical_test_skyshi/activity/controller"
	mysqlActivity "technical_test_skyshi/activity/repository/mysql"
	activityService "technical_test_skyshi/activity/service"

	todoController "technical_test_skyshi/todo/controller"
	mysqlTodo "technical_test_skyshi/todo/repository/mysql"
	todoService "technical_test_skyshi/todo/service"

	"technical_test_skyshi/database"
	"technical_test_skyshi/router"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(
		database.InitDBMysql,

		mysqlActivity.NewMysqlAcitivity,
		activityService.NewActivityServiceImpl,
		activityController.NewActivityControllerImpl,

		mysqlTodo.NewMysqlTodo,
		todoService.NewTodoServiceImpl,
		todoController.NewTodoControllerImpl,

		router.SetupRouter,
	)

	return nil
}
