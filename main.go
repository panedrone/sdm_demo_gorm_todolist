package main

import (
	"github.com/gin-gonic/gin"
	"log"
	h "sdm_demo_gorm_todolist/api_handlers"
	"sdm_demo_gorm_todolist/dbal"
)

func main() {
	err := dbal.OpenDB()
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		_ = dbal.CloseDB()
	}()
	// SetMode before getting gin.New()
	gin.SetMode(gin.ReleaseMode)
	// use New instead of Default to avoid HTTP logging
	engine := gin.New()
	// https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/
	// === panedrone: type "http://localhost:8080/assets/" to render index.html
	engine.Static("/assets", "./assets")
	{
		groups := engine.Group("/groups")
		groups.GET("/", h.GroupsReadAllHandler)
		groups.POST("/", h.GroupCreateHandler)
		{
			group := groups.Group("/:g_id")
			group.GET("/", h.GroupReadHandler)
			group.PUT("/", h.GroupUpdateHandler)
			group.DELETE("/", h.GroupDeleteHandler)
			{
				groupTasks := group.Group("/tasks")
				groupTasks.GET("/", h.TasksByGroupReadHandler)
				groupTasks.POST("/", h.TaskCreateHandler)
			}
		}
	}
	{
		task := engine.Group("/tasks/:t_id")
		task.GET("", h.TaskReadHandler)
		task.PUT("", h.TaskUpdateHandler)
		task.DELETE("", h.TaskDeleteHandler)
	}
	log.Fatal(engine.Run(":8080"))
}
