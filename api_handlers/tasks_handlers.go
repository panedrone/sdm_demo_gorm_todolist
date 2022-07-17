package api_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_gorm_todolist/dal"
	"sdm_demo_gorm_todolist/dal/dao"
	"time"
)

func ReturnTaskHandler(ctx *gin.Context) {
	var inTsk taskUri
	err := ctx.ShouldBindUri(&inTsk)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	currTask := dal.Task{TId: inTsk.TId}
	err = dao.Db().Take(&currTask).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, currTask)
}

func ReturnGroupTasksHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	// don't fetch comments for list:
	var tasks []dal.Task // https://gorm.io/docs/query.html
	err := dao.Db().Table("tasks").Where("g_id = ?", uri.GId).Order("t_id").
		Select("t_id", "t_date", "t_subject", "t_priority").Find(&tasks).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func TaskCreateHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	var inTask dal.Task
	err := ctx.ShouldBindJSON(&inTask)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	t := dal.Task{}
	t.GId = uri.GId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	currentTime := time.Now().Local()
	layoutISO := currentTime.Format("2006-01-02")
	t.TDate = layoutISO
	err = dao.Db().Create(&t).Error
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
	currTask := dal.Task{TId: inTsk.TId}
	err = dao.Db().Delete(&currTask).Error
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
	var inTask dal.Task
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
	t := dal.Task{TId: inTsk.TId}
	err = dao.Db().Take(&t).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	t.TSubject = inTask.TSubject
	t.TPriority = inTask.TPriority
	t.TDate = inTask.TDate
	t.TComments = inTask.TComments
	err = dao.Db().Save(&t).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}
