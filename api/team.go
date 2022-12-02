package api

import (
	"SK-todo/operation"
	"github.com/gin-gonic/gin"
)

// CreateTeam @Tags TEAM
//
//	@Summary	Create team
//	@Produce	json
//	@Accept		json
//	@Header		200		{string}	Authorization			"required"
//	@Param		data	body		operation.CreateTeamOP	true	"team name"
//	@Success	200		{object}	serializer.ResponseTeam	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTeam	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/team [post]
func CreateTeam(c *gin.Context) {
	var createService operation.CreateTeamOP
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.CreateTeam()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateTeam @Tags TEAM
//
//	@Summary	Update team info
//	@Produce	json
//	@Accept		json
//
//	@param		id		path		int						true	"team id"
//
//	@Header		200		{string}	Authorization			"required"
//	@Param		data	body		operation.UpdateTeamOP	true	"team info"
//	@Success	200		{object}	serializer.ResponseTeam	"{"status":200,"data":{},"msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseTeam	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/:team_id [put]
func UpdateTeam(c *gin.Context) {
	var updateService operation.UpdateTeamOP
	if err := c.ShouldBind(&updateService); err == nil {
		res := updateService.UpdateTeam(c.Param("team_id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteTeam @Tags TEAM
//
//	@Summary	Delete team
//	@Produce	json
//	@Accept		json
//
//	@param		id	path		int						true	"team id"
//
//	@Header		200	{string}	Authorization			"required"
//	@Success	200	{object}	serializer.ResponseTeam	"{"status":200,"data":{},"msg":"ok"}"
//	@Failure	500	{object}	serializer.ResponseTeam	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/:team_id [delete]
func DeleteTeam(c *gin.Context) {
	var deleteService operation.DeleteTeamOP
	res := deleteService.DeleteTeam(c.Param("team_id"))
	c.JSON(200, res)
}

// ShowTeamMember @Tags TEAM
//
//	@Summary	Show team members
//	@Produce	json
//	@Accept		json
//
//	@param		id	path		int								true	"team id"
//
//	@Header		200	{string}	Authorization					"required"
//	@Success	200	{object}	serializer.ResponseTeamMembers	"{"status":200,"data":{},"msg":"ok"}"
//	@Failure	500	{object}	serializer.ResponseTeamMembers	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/team/:team_id [get]
func ShowTeamMember(c *gin.Context) {
	var showService operation.ShowTeamMemberOP
	res := showService.ShowTeamMember(c.Param("team_id"))
	c.JSON(200, res)
}
