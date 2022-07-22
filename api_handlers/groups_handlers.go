package api_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_gorm_todolist/dbal"
	"sdm_demo_gorm_todolist/models"
)

func GroupCreateHandler(ctx *gin.Context) {
	var inGr newGroup
	err := ctx.ShouldBindJSON(&inGr)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	gr := models.Group{}
	gr.GName = inGr.GName
	err = dbal.Db().Create(&gr).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}

func ReturnAllGroupsHandler(ctx *gin.Context) {
	grDao := dbal.NewGroupsDao()
	groups, err := grDao.GetAllGroupsEx()
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	respondWithJSON(ctx, http.StatusOK, groups)
}

func GroupUpdateHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.BindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	var inGroup models.Group
	err := ctx.ShouldBindJSON(&inGroup)
	if err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	gr := models.Group{GId: uri.GId}
	err = dbal.Db().Take(&gr).Error
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	gr.GName = inGroup.GName
	err = dbal.Db().Save(&gr).Error // https://gorm.io/docs/update.html
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
	gr := models.Group{GId: uri.GId}
	err := dbal.Db().Delete(&gr).Error // https://gorm.io/docs/delete.html
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
	group := models.Group{GId: uri.GId}
	err := dbal.Db().Take(&group).Error
	if err == gorm.ErrRecordNotFound {
		respondWithNotFoundError(ctx, err.Error())
	} else if err != nil {
		respondWith500(ctx, err.Error())
	} else {
		respondWithJSON(ctx, http.StatusOK, group)
	}
}
