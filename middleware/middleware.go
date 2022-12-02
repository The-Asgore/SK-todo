package middleware

import (
	"SK-todo/auth"
	"SK-todo/model"
	respCode "SK-todo/resp-code"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims *auth.Claims
		var err error
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 30004
		} else {
			claims, err = auth.ParseToken(token)
			if err != nil {
				code = respCode.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = respCode.ErrorAuthCheckTokenTimeout
			}
		}
		if code != respCode.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    respCode.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Set("claimsID", claims.Id)
		c.Next()
	}
}

func IsUserHasTeamAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var team model.Team
		teamID := c.Param("team_id")
		if teamID == "" {
			teamID = c.DefaultQuery("team_id", "0")
		}
		model.DB.Preload("User").First(&team, teamID)
		for _, member := range team.User {
			claimID, _ := c.Get("claimsID")
			if member.ID == claimID.(uint) {
				c.Next()
				return
			}
		}
		code := respCode.ErrorAuthCheckTokenFail
		c.JSON(400, gin.H{
			"status": code,
			"msg":    respCode.GetMsg(code),
			"data":   data,
		})
		c.Abort()
		return
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,"+
				"session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, "+
				"X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, "+
				"Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
