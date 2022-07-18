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

func ReturnTaskHandler(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	task := models.Task{TId: inTsk.TId}
	err = dbal.Db().Take(&task).Error
	if err == gorm.ErrRecordNotFound {
		respondWithNotFoundError(ctx, err.Error())
	} else if err != nil {
		respondWith500(ctx, err.Error())
	} else {
		respondWithJSON(ctx, http.StatusOK, task)
	}
}

func ReturnGroupTasksHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	// don't fetch comments for list:
	var tasks []models.Task // https://gorm.io/docs/query.html
	err := dbal.Db().Table("tasks").Where("g_id = ?", uri.GId).Order("t_id").
		Select("t_id", "t_date", "t_subject", "t_priority").Find(&tasks).Error
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
	err = dbal.Db().Create(&t).Error
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
	currTask := models.Task{TId: inTsk.TId}
	err = dbal.Db().Delete(&currTask).Error
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
	t := models.Task{TId: inTsk.TId}
	err = dbal.Db().Take(&t).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	err = dbal.Db().Save(&t).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}
