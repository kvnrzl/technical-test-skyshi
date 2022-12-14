package helper

import (
	"fmt"
	"net/http"
	"technical_test_skyshi/model"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, model.ResponseError{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseErrRecordNotFound(c *gin.Context, id int) {
	c.JSON(http.StatusNotFound, model.ResponseError{
		Status:  http.StatusText(http.StatusNotFound),
		Message: fmt.Errorf("Activity with ID %d not found", id).Error(),
		Data:    struct{}{},
	})
}

func ResponseErrTitleRequired(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, model.ResponseError{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseInternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, model.ResponseError{
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: message,
		Data:    struct{}{},
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.ResponseStatusOK{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}
