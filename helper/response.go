package helper

import (
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

func ResponseRecordNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, model.ResponseError{
		Status:  http.StatusText(http.StatusNotFound),
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

func ResponseStatusOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.ResponseStatusOK{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}

func ResponseStatusCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, model.ResponseStatusOK{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}
