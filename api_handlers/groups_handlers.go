package api_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_go_todolist/dal"
)

func GroupCreateHandler(ctx *gin.Context) {
	var inGr newGroup
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	gr := dal.Group{}
	gr.GName = inGr.GName
	err = dal.Db().Create(&gr).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}

func ReturnAllGroupsHandler(ctx *gin.Context) {
	dao := dal.NewGroupsDao()
	groups, err := dao.GetGroupsEx()
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, groups)
}

func GroupUpdateHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.BindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	var inGroup dal.Group
	err := ctx.ShouldBindJSON(&inGroup)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	gr := dal.Group{GId: uri.GId}
	err = dal.Db().First(&gr).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	gr.GName = inGroup.GName
	err = dal.Db().Save(&gr).Error // https://gorm.io/docs/update.html
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}

func GroupDeleteHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.BindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	gr := dal.Group{GId: uri.GId}
	err := dal.Db().Delete(&gr).Error // https://gorm.io/docs/delete.html
	if err != nil {
		respondWith500(ctx, err.Error())
	}
}

func ReturnGroupHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	gr := dal.Group{GId: uri.GId}
	err := dal.Db().First(&gr).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gr)
}
