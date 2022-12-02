package api

import (
	"SK-todo/operation"
	"github.com/gin-gonic/gin"
)

// UserRegister @Tags USER
//
//	@Summary	User register
//	@Produce	json
//	@Accept		json
//	@Param		data	body		operation.UserOP		true	"username, password"
//	@Success	200		{object}	serializer.ResponseUser	"{"status":200,"data":{},"msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseUser	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var userRegisterService operation.UserOP
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// UserLogin @Tags USER
//
//	@Summary	User login
//	@Produce	json
//	@Accept		json
//	@Param		data	body		operation.UserOP		true	"user_name, password"
//	@Success	200		{object}	serializer.ResponseUser	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseUser	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginService operation.UserOP
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// UserShowTeams @Tags USER
//
//	@Summary	Show user's teams
//	@Produce	json
//	@Accept		json
//	@Success	200	{object}	serializer.ResponseUserTeams	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500	{object}	serializer.ResponseUserTeams	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/teams [get]
func UserShowTeams(c *gin.Context) {
	var userLoginService operation.ShowUserTeamOP
	chaimID, _ := c.Get("claimsID")
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.ShowTeams(chaimID.(uint))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// UpdateUser @Tags USER
//
//	@Summary	Update user info
//	@Produce	json
//	@Accept		json
//
//	@Param		data	body		operation.UpdateUserOP			true	"username, password"
//
//	@Success	200		{object}	serializer.ResponseUserTeams	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseUserTeams	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user [put]
func UpdateUser(c *gin.Context) {
	var userLoginService operation.UpdateUserOP
	chaimID, _ := c.Get("claimsID")
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.UpdateUser(chaimID.(uint))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// DeleteUser @Tags USER
//
//	@Summary	Update user info
//	@Produce	json
//	@Accept		json
//
//	@pragma		uid																																																	path								string				true	"uid"
//
//	@Success	200	{object}	serializer.ResponseUserTeams	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500	{object}	serializer.ResponseUserTeams	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user [delete]
func DeleteUser(c *gin.Context) {
	var userLoginService operation.DeleteUserOP
	chaimID, _ := c.Get("claimsID")
	res := userLoginService.DeleteUser(chaimID.(uint))
	c.JSON(200, res)
}

// JoinTeam @Tags USER
//
//	@Summary	Join team
//	@Produce	json
//	@Accept		json
//
//	@Param		team_id	path		string							true	"team_id"
//
//	@Success	200		{object}	serializer.ResponseUserTeams	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseUserTeams	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/team/:team_id [put]
func JoinTeam(c *gin.Context) {
	var userLoginService operation.JoinTeamOP
	chaimID, _ := c.Get("claimsID")
	res := userLoginService.JoinTeam(chaimID.(uint), c.Param("team_id"))
	c.JSON(200, res)
}

// LeaveTeam @Tags USER
//
//	@Summary	Leave team
//	@Produce	json
//	@Accept		json
//
//	@Param		team_id	path		string							true	"team_id"
//
//	@Success	200		{object}	serializer.ResponseUserTeams	"{"Status":true,"data":{},"Msg":"ok"}"
//	@Failure	500		{object}	serializer.ResponseUserTeams	"{"status":500,"data":{},"Msg":{},"Error":"error"}"
//	@Router		/api/v1/user/team/:team_id [delete]
func LeaveTeam(c *gin.Context) {
	var userLoginService operation.LeaveTeamOP
	chaimID, _ := c.Get("claimsID")
	res := userLoginService.LeaveTeam(chaimID.(uint), c.Param("team_id"))
	c.JSON(200, res)
}
