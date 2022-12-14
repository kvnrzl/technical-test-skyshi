package router

import (
	activityController "technical_test_skyshi/activity/controller"
	todoController "technical_test_skyshi/todo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	activityController activityController.ActivityController,
	todoController todoController.TodoController,
) *gin.Engine {
	router := gin.Default()

	router.Use(Middleware())

	router.GET("/", func(r *gin.Context) {
		r.JSON(200, "API service is ready!")
	})

	activity := router.Group("/activity-groups")
	{
		activity.POST("/", activityController.Create)
		activity.GET("/", activityController.GetAll)
		activity.GET("/:id", activityController.GetByID)
		activity.PATCH("/:id", activityController.Update)
		activity.DELETE("/:id", activityController.Delete)
	}

	todo := router.Group("/todo-items")
	{
		todo.POST("/", todoController.Create)
		todo.GET("/", todoController.GetAll)
		todo.GET("/:id", todoController.GetByID)
		todo.PATCH("/:id", todoController.Update)
		todo.DELETE("/:id", todoController.Delete)
	}

	return router
}
