package app

import (
	"technical_test_skyshi/activity/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(activityController controller.ActivityController) *gin.Engine {
	router := gin.Default()
	// router.Use(CORSMiddleware())

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

	return router
}
