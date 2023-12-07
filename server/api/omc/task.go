package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/heypkg/store/jsontype"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type listTasksData struct {
	Data  []omc.Task `json:"Data"`
	Total int64      `json:"Total"`
}

// HandleListTasks lists all tasks.
// @Summary List tasks
// @ID list-tasks
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listTasksData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/tasks [get]
// @Tags OMC Tasks
func HandleListTasks(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.Task](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listTasksData{Data: data, Total: total})
}

type createTaskBody struct {
	TaskName       string            `json:"TaskName"`
	TaskType       string            `json:"TaskType"`
	TaskStatus     omc.TaskStatus    `json:"TaskStatus"`
	ExecMode       omc.TaskExecMode  `json:"ExecMode"`
	ExecStartTime  jsontype.JSONTime `json:"ExecStartTime"`
	ExecEndTime    jsontype.JSONTime `json:"ExecEndTime"`
	ExecInterval   int               `json:"ExecInterval"`
	ExecTimes      int               `json:"ExecTimes"`
	ExecProcess    int               `json:"ExecProcess"`
	Creater        string            `json:"Creater"`
	MetaData       *jsontype.Tags    `json:"MetaData"`
	DeviceIds      []uint            `json:"DeviceIds"`
	ParameterNames []string          `json:"ParameterNames"`
}

// HandleCreateTask creates a new task.
// @Summary Create task
// @ID create-task
// @Produce json
// @Security Bearer
// @Param body body createTaskBody true "Task"
// @Success 200 {object} omc.Task
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/tasks [post]
// @Tags OMC Tasks
func HandleCreateTask(c echo.Context) error {
	var data createTaskBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	ranges := jsontype.Tags{}
	if len(data.DeviceIds) > 0 {
		ranges["DeviceIds"] = data.DeviceIds
	}
	if data.ParameterNames != nil {
		ranges["ParameterNames"] = data.ParameterNames
	}

	db := store.GetDB()
	task := omc.Task{
		Schema:        cast.ToString(c.Get("schema")),
		TaskName:      data.TaskName,
		TaskType:      data.TaskType,
		TaskStatus:    data.TaskStatus,
		ExecMode:      data.ExecMode,
		ExecStartTime: data.ExecStartTime,
		ExecEndTime:   data.ExecEndTime,
		ExecInterval:  data.ExecInterval,
		ExecTimes:     data.ExecTimes,
		ExecProcess:   data.ExecProcess,
		Creater:       data.Creater,
		MetaData:      data.MetaData,
		Ranges:        &ranges,
	}

	result := db.Create(&task)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	tsrv := omc.GetTaskServer()
	tsrv.StartTask(&task)
	return c.JSON(http.StatusOK, task)
}

// HandleGetTask retrieves a single task.
// @Summary Get task
// @ID get-task
// @Produce json
// @Security Bearer
// @Param id path int true "Task ID"
// @Success 200 {object} omc.Task
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/tasks/{id} [get]
// @Tags OMC Tasks
func HandleGetTask(c echo.Context) error {
	task := echohandler.GetObjectFromEchoContext[omc.Task](c)
	return c.JSON(http.StatusOK, task)
}

// HandleDeleteTask deletes the task.
// @Summary Delete task
// @Tags OMC Tasks
// @ID delete-task
// @Security Bearer
// @Param id path int true "Task ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/tasks/{id} [delete]
func HandleDeleteTask(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.Task](c)
	db := store.GetDB().Unscoped()
	if result := db.Delete(obj); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	tsrv := omc.GetTaskServer()
	tsrv.ClearTask(obj)
	return c.NoContent(http.StatusNoContent)
}
