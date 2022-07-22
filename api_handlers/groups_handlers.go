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
	err = dbal.NewGroupsDao().Create(&gr)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
}

func GroupsReadAllHandler(ctx *gin.Context) {
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
	dao := dbal.NewGroupsDao()
	gr, err := dao.Read(uri.GId)
	if err != nil {
		respondWith500(ctx, err.Error())
		return
	}
	gr.GName = inGroup.GName
	_, err = dao.Update(gr)
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
	gr := &models.Group{
		GId: uri.GId,
	}
	_, err := dbal.NewGroupsDao().Delete(gr)
	if err != nil {
		respondWith500(ctx, err.Error())
	}
}

func GroupReadHandler(ctx *gin.Context) {
	var uri groupUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		respondWithBadRequestError(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
		return
	}
	group, err := dbal.NewGroupsDao().Read(uri.GId)
	if err == gorm.ErrRecordNotFound {
		respondWithNotFoundError(ctx, err.Error())
	} else if err != nil {
		respondWith500(ctx, err.Error())
	} else {
		respondWithJSON(ctx, http.StatusOK, group)
	}
}
