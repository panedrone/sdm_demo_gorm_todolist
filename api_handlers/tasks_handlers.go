package api_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_gorm_todolist/dbal"
	"sdm_demo_gorm_todolist/models"
	"time"
)

func TaskReadHandler(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	task, err := dbal.NewTasksDao().ReadTask(inTsk.TId)
	if err == gorm.ErrRecordNotFound {
		respondWithNotFoundError(ctx, err.Error())
	} else if err != nil {
		respondWith500(ctx, err.Error())
	} else {
		respondWithJSON(ctx, http.StatusOK, task)
	}
}

func TasksByGroupReadHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	var tasks []*models.Task
	tasks, err := dbal.NewTasksDao().ReadGroupTasks(uri.GId)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, tasks)
}

func TaskCreateHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	var inTask models.Task
	err := ctx.ShouldBindJSON(&inTask)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	t := models.Task{}
	t.GId = uri.GId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	currentTime := time.Now().Local()
	layoutISO := currentTime.Format("2006-01-02")
	t.TDate = layoutISO
	err = dbal.NewTasksDao().CreateTask(&t)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}

func TaskDeleteHandler(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	tsk := &models.Task{
		TId: inTsk.TId,
	}
	_, err = dbal.NewTasksDao().DeleteTask(tsk)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}

func TaskUpdateHandler(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	var inTask models.Task
	err = ctx.ShouldBindJSON(&inTask)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	_, err = time.Parse("2006-01-02", inTask.TDate)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid input: %s", err.Error()))
		return
	}
	if len(inTask.TSubject) == 0 {
		respondWithBadRequestError(ctx, fmt.Sprintf("Subject required"))
		return
	}
	if inTask.TPriority <= 0 {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid Priority: %d", inTask.TPriority))
		return
	}
	dao := dbal.NewTasksDao()
	t, err := dao.ReadTask(inTsk.TId)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	_, err = dao.UpdateTask(t)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}
