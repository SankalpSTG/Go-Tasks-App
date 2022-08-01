package handler

import (
	"goapp/dao"
	error "goapp/misc/error"
	misc "goapp/misc/error"
	"goapp/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	var isCompleted *bool
	parsedBool, err := strconv.ParseBool(queryParams.Get("is_completed"))
	if(err != nil){ isCompleted = nil}else{isCompleted = &parsedBool}
	tasks, _ := repository.GetTasks(isCompleted)
	c.JSON(http.StatusOK, misc.TasksData{Data: tasks})
}
func InsertTask(c *gin.Context) {
	task := new(dao.Task)
	err := c.Bind(task)
	if(err == nil){
		task.CreatedAt = time.Now().Unix()
		SerialID, _ := repository.InsertTask(task)
		c.JSON(http.StatusOK, misc.IntData{Data: SerialID})
		return
	}
	c.JSON(http.StatusInternalServerError, error.Error{Error: true, Message: "Couldn't Add Task"})
}
func UpdateTask(c *gin.Context){
	task := new(dao.Task)
	err := c.Bind(task)
	if(err == nil){
		task.UpdatedAt = time.Now().Unix()
		err := repository.UpdateTask(task)
		if(err != nil){
			c.JSON(http.StatusInternalServerError, misc.Error{Error: true, Message: "Couldn't Update Task"})
			return
		}
		c.JSON(http.StatusOK, misc.Error{Error: false, Message: "Updated Task"})
		return
	}
	c.JSON(http.StatusInternalServerError, error.Error{Error: true, Message: "Couldn't Add Task"})
}