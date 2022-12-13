package helper

import (
	"fmt"
	"net/http"
	"technical_test_skyshi/model"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, model.Response{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
		Data:    nil,
	})
}

func ResponseErrRecordNotFound(c *gin.Context, id int) {
	c.JSON(http.StatusNotFound, model.Response{
		Status:  http.StatusText(http.StatusNotFound),
		Message: fmt.Errorf("Activity with ID %d not found", id).Error(),
		Data:    nil,
	})
}

func ResponseErrTitleRequired(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, model.Response{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
		Data:    nil,
	})
}

func ResponseInternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, model.Response{
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: message,
		Data:    nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.Response{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}
