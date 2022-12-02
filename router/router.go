package router

import (
	"SK-todo/api"
	_ "SK-todo/docs"
	"SK-todo/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("sk-secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(sessions.Sessions("sk-sessions", store))
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())

		userOP := authed.Group("user")
		{
			userOP.GET("tasks", api.ListPersonalTasks)
			userOP.POST("task", api.CreatePersonalTask)
			userOP.GET("task/:task_id", api.ShowPersonalTask)
			userOP.DELETE("task/:task_id", api.DeletePersonalTask)
			userOP.PUT("task/:task_id", api.UpdatePersonalTask)
			userOP.POST("search", api.SearchPersonalTasks)

			userOP.GET("teams", api.UserShowTeams)
			userOP.POST("team", api.CreateTeam)
			userOP.PUT("team/:team_id", api.JoinTeam)
			userOP.DELETE("team/:team_id", api.LeaveTeam)
			userOP.PUT("/", api.UpdateUser)
			userOP.DELETE("/", api.DeleteUser)

		}

		teamOP := authed.Group("team")
		teamOP.Use(middleware.IsUserHasTeamAuth())
		{
			teamOP.PUT("/:team_id", api.UpdateTeam)
			teamOP.DELETE("/:team_id", api.DeleteTeam)
			teamOP.GET("/:team_id", api.ShowTeamMember)

			teamOP.POST("task", api.CreateTeamTask)
			teamOP.GET("tasks", api.ListTeamTasks)
			teamOP.GET("task/:task_id", api.ShowTeamTask)
			teamOP.DELETE("task/:task_id", api.DeleteTeamTask)
			teamOP.PUT("task/:task_id", api.UpdateTeamTask)
			teamOP.POST("search", api.SearchTeamTasks)
		}
	}
	return r
}
