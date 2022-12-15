package controller

import (
	"strconv"
	"technical_test_skyshi/helper"
	"technical_test_skyshi/model"
	"technical_test_skyshi/todo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoControllerImpl struct {
	todoService service.TodoService
}

func NewTodoControllerImpl(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		todoService: todoService,
	}
}

func (t *TodoControllerImpl) Create(c *gin.Context) {
	todoCreate := &model.TodoCreate{}
	c.ShouldBindJSON(&todoCreate)

	todo, err := t.todoService.Create(c, todoCreate)
	if err == model.ErrTitleRequired {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	if err == model.ErrActivityGroupIDRequired {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusCreated(c, todo)
}

func (t *TodoControllerImpl) GetAll(c *gin.Context) {
	query := c.Query("activity_group_id")
	todos, err := t.todoService.GetAll(c, query)
	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, todos)
}

func (t *TodoControllerImpl) GetByID(c *gin.Context) {
	id := c.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	parseID64 := int64(parseID)
	todo, err := t.todoService.GetByID(c, parseID64)
	if err == gorm.ErrRecordNotFound {
		helper.ResponseTodoNotFound(c, parseID64)
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, todo)
}

func (t *TodoControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	todo := &model.Todo{}
	c.ShouldBindJSON(&todo)
	todo.ID = int64(parseID)

	todo, err = t.todoService.Update(c, todo)
	if err == gorm.ErrRecordNotFound {
		helper.ResponseTodoNotFound(c, int64(parseID))
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, todo)
}

func (t *TodoControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	parseID, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	parseID64 := int64(parseID)
	err = t.todoService.Delete(c, parseID64)
	if err == gorm.ErrRecordNotFound {
		helper.ResponseTodoNotFound(c, parseID64)
		return
	}

	if err != nil {
		helper.ResponseInternalServerError(c, err.Error())
		return
	}

	helper.ResponseStatusOK(c, struct{}{})
}
