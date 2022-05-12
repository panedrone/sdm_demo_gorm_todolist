package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sdm_demo_go_todolist/api_handlers"
	"sdm_demo_go_todolist/dal"
)

func main() {
	err := dal.OpenDB()
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		err = dal.CloseDB()
	}()

	gin.SetMode(gin.ReleaseMode)
	myRouter := gin.Default()
	// https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/
	// === panedrone: type "http://localhost:8080/assets/" to render index.html
	myRouter.Static("/assets", "./assets")
	////////////////////
	{
		groups := myRouter.Group("/groups")
		groups.GET("/", api_handlers.ReturnAllGroupsHandler)
		groups.POST("/", api_handlers.GroupCreateHandler)
		{
			group := groups.Group("/:g_id")
			group.GET("/", api_handlers.ReturnGroupHandler)
			group.PUT("/", api_handlers.GroupUpdateHandler)
			group.DELETE("/", api_handlers.GroupDeleteHandler)
			{
				tasks := group.Group("/tasks")
				tasks.GET("/", api_handlers.ReturnGroupTasksHandler)
				tasks.POST("/", api_handlers.TaskCreateHandler)
			}
		}
	}
	////////////////////
	{
		task := myRouter.Group("/tasks/:t_id")
		task.GET("", api_handlers.ReturnTaskHandler)
		task.PUT("", api_handlers.TaskUpdateHandler)
		task.DELETE("", api_handlers.TaskDeleteHandler)
	}
	// log.Fatal
	// https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
	// log.Fatal(http.ListenAndServe(":8080", myRouter))
	// https://stackoverflow.com/questions/57354389/how-to-render-static-files-within-gin-router
	log.Fatal(myRouter.Run(":8080"))
}
