package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../model"
)

const (
	HOST = "0.0.0.0:1234"
)

func Serve() error {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	user := router.Group("/user")
	{
		router.GET("/", GetUserList)
	}

	{
		user.POST("/add", AddUser)
		user.GET("/detail", GetUserDetail)
		user.POST("/update", UpdateUser)
	}

	return router.Run(HOST)
}

func GetUserList(c *gin.Context) {
	//users := []*model.Users{}
	//for id, user := range users {
	//
	//}
}

func GetUserDetail(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}