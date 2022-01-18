package handlers

import (
	"finalproject/database"
	"finalproject/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/swag/example/celler/httputil"
)

type APIEnv struct {
	DB *gorm.DB
}

//GetTodo godoc
//@summary Get Todo
//@Description get All todo
//@Tags todos
//@Produce json
//@Success 200 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
//@Router /todos [get]
func (a *APIEnv) GetTodos(c *gin.Context) {
	todos, err := database.GetTodos(a.DB)
	if err != nil {
		fmt.Println(err)
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetTodo godoc
// @Tags todos
// @Description Get todo
// @ID get-todo
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos/{id} [get]
func (a *APIEnv) GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	todo, exists, err := database.GetTodoByID(i, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	if !exists {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// CreateTodo godoc
// @Tags todos
// @Description Create todo
// @ID create-todo
// @Accept json
// @Produce json
// @Param RequestBody body models.ReqTodo true "request body json"
// @Success 201 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos [post]
func (a *APIEnv) CreateTodo(c *gin.Context) {
	todo := models.Todo{}
	err := c.BindJSON(&todo)
	if err != nil {

		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if err := a.DB.Create(&todo).Error; err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Tags todos
// @Description Delete todo
// @ID delete-todo
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos/{id} [delete]
func (a *APIEnv) DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	_, exists, err := database.GetTodoByID(i, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	err = database.DeleteTodo(id, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}

// UpdateTodo godoc
// @Tags todos
// @Description update todo
// @ID update-todo
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param RequestBody body models.Req	Todo true "request body json"
// @Success 200 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos/{id} [put]
func (a *APIEnv) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}
	_, exists, err := database.GetTodoByID(i, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	updatedTodo := models.Todo{}
	err = c.BindJSON(&updatedTodo)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	updatedTodo.ID = uint(i)

	if err := database.UpdateTodo(a.DB, &updatedTodo); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	a.GetTodo(c)
}
