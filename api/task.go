package api

import (
	"SK-todo/operation"
	"github.com/gin-gonic/gin"
)

// CreatePersonalTask @Tags TASK
//
//	@Summary	Create Personal Task
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//	@Param		data	body		operation.CreateTaskOP	true	"task info"
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/task [post]
func CreatePersonalTask(c *gin.Context) {
	createService := operation.CreateTaskOP{}
	chaimID, _ := c.Get("claimsID")
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.CreatePersonalTask(chaimID.(uint))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// CreateTeamTask @Tags TASK
//
//	@Summary	Create Team Task
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//	@Param		data	body		operation.CreateTaskOP	true	"task info"
//
//	@Param		team_id	query		uint					true	"team id is required when task is a team task"
//
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/task [post]
func CreateTeamTask(c *gin.Context) {
	createService := operation.CreateTaskOP{}
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.CreateTeamTask(c.DefaultQuery("team_id", "0"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// ListPersonalTasks @Tags TASK
//
//	@Summary	Get Personal task list
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//	@Param		Index	query		string					false	"create_at | id | due_date | status"
//	@Param		order	query		string					false	"asc | desc"
//	@Param		limit	query		uint					false	"recode limit"
//	@Param		start	query		uint					false	"start from"
//
//	@param		status	query		string					false	"status"
//
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/tasks [get]
func ListPersonalTasks(c *gin.Context) {
	listService := operation.ListTasksOP{}
	listService.Index = c.DefaultQuery("index", "created_at")
	listService.Order = c.DefaultQuery("order", "desc")
	listService.Status = c.Query("status")
	if listService.Order != "desc" && listService.Order != "asc" {
		listService.Order = "desc"
	}
	listService.Limit = c.DefaultQuery("limit", "15")
	listService.Start = c.DefaultQuery("start", "0")
	chaimID, _ := c.Get("claimsID")
	res := listService.ListPersonalTask(chaimID.(uint))
	c.JSON(200, res)
}

// ListTeamTasks @Tags TASK
//
//	@Summary	Get Team task list
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//	@Param		Index	query		string					false	"created_at | id | due_date | status"
//	@Param		order	query		string					false	"asc | desc"
//	@Param		limit	query		uint					false	"record limit"
//	@Param		start	query		uint					false	"start from"
//	@Param		team_id	query		uint					true	"team id"
//
//	@param		status	query		string					false	"status"
//
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/tasks [get]
func ListTeamTasks(c *gin.Context) {
	listService := operation.ListTasksOP{}
	listService.Index = c.DefaultQuery("index", "created_at")
	listService.Order = c.DefaultQuery("order", "desc")
	listService.Status = c.Query("status")
	if listService.Order != "desc" && listService.Order != "asc" {
		listService.Order = "desc"
	}
	listService.Limit = c.DefaultQuery("limit", "15")
	listService.Start = c.DefaultQuery("start", "0")
	teamID := c.DefaultQuery("team_id", "0")
	res := listService.ListTeamTask(teamID)
	c.JSON(200, res)

}

// ShowPersonalTask @Tags TASK
//
//	@Summary	Show detailed information of personal task
//	@Produce	json
//	@Accept		json
//	@Header		200	{string}	Authorization			"required"
//
//	@param		id	path		string					true	"task id"
//
//	@Success	200	{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500	{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/task/:task_id [get]
func ShowPersonalTask(c *gin.Context) {
	chaimID, _ := c.Get("claimsID")
	showTaskService := operation.ShowTaskOP{UserID: chaimID.(uint)}
	res := showTaskService.ShowPersonalTask(c.Param("task_id"))
	c.JSON(200, res)
}

// ShowTeamTask @Tags TASK
//
//	@Summary	Show detailed information of team task
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//
//	@param		id		path		string					true	"task id"
//
//	@param		team_id	query		string					true	"team id"
//
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/task/:task_id [get]
func ShowTeamTask(c *gin.Context) {
	showTaskService := operation.ShowTaskOP{TeamID: c.DefaultQuery("team_id", "0")}
	res := showTaskService.ShowTeamTask(c.Param("task_id"))
	c.JSON(200, res)
}

// DeletePersonalTask @Tags TASK
//
//	@Summary	Delete personal task
//	@Produce	json
//	@Accept		json
//
//	@param		id	path		string					true	"task id"
//
//	@Header		200	{string}	Authorization			"required"
//	@Success	200	{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500	{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/user/task/:task_id [delete]
func DeletePersonalTask(c *gin.Context) {
	chaimID, _ := c.Get("claimsID")
	deleteTaskService := operation.DeleteTaskOP{}
	res := deleteTaskService.DeletePersonalTask(c.Param("task_id"), chaimID.(uint))
	c.JSON(200, res)
}

// DeleteTeamTask @Tags TASK
//
//	@Summary	Delete team task
//	@Produce	json
//	@Accept		json
//
//	@param		id		path		string					true	"task id"
//
//	@Header		200		{string}	Authorization			"required"
//	@Param		team_id	query		uint					true	"team id"
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/task/:task_id [delete]
func DeleteTeamTask(c *gin.Context) {
	deleteTaskService := operation.DeleteTaskOP{}
	teamID := c.DefaultQuery("team_id", "0")
	res := deleteTaskService.DeleteTeamTask(c.Param("task_id"), teamID)
	c.JSON(200, res)
}

// UpdatePersonalTask @Tags TASK
//
//	@Summary	Update personal task
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//
//	@param		id		path		string					true	"task id"
//
//	@Param		data	body		operation.UpdateTaskOP	true	"task info"
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/task/:task_id [put]
func UpdatePersonalTask(c *gin.Context) {
	updateTaskService := operation.UpdateTaskOP{}
	chaimID, _ := c.Get("claimsID")
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.UpdatePersonalTask(c.Param("task_id"), chaimID.(uint))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// UpdateTeamTask @Tags TASK
//
//	@Summary	Update team task
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//
//	@param		id		path		string					true	"task id"
//
//	@Param		data	body		operation.UpdateTaskOP	true	"task info"
//	@Param		team_id	query		uint					true	"team id"
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/task/:task_id [put]
func UpdateTeamTask(c *gin.Context) {
	updateTaskService := operation.UpdateTaskOP{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.UpdateTeamTask(c.Param("task_id"), c.DefaultQuery("team_id", "0"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// SearchPersonalTasks @Tags TASK
//
//	@Summary	Search for personal tasks
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization
//	@Header		200		{string}	Authorization			"required"
//	@Param		data	body		operation.SearchTaskOP	true	"search task info"
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/search [post]
func SearchPersonalTasks(c *gin.Context) {
	searchTaskService := operation.SearchTaskOP{}
	chaimID, _ := c.Get("claimsID")
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.SearchPersonalTask(chaimID.(uint))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// SearchTeamTasks @Tags TASK
//
//	@Summary	Search for team tasks
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization
//	@Header		200		{string}	Authorization			"required"
//	@Param		data	body		operation.SearchTaskOP	true	"search task info"
//	@Param		team_id	query		uint					true	"team id"
//	@Success	200		{object}	serializer.ResponseTask	"{"Status":200,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTask	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/search [post]
func SearchTeamTasks(c *gin.Context) {
	searchTaskService := operation.SearchTaskOP{}
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.SearchTeamTask(c.DefaultQuery("team_id", "0"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
