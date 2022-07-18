package api_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type groupUri struct {
	GId int64 `uri:"g_id" binding:"required"`
}

type newGroup struct {
	GName string `json:"g_name" binding:"required"`
}

type taskUri struct {
	TId int64 `uri:"t_id" binding:"required"`
}

type Err struct {
	Error string `json:"error"`
}

func respondWithBadRequestError(ctx *gin.Context, message string) {
	respondWithError(ctx, http.StatusBadRequest, message)
}

func respondWithNotFoundError(ctx *gin.Context, message string) {
	respondWithError(ctx, http.StatusNotFound, message)
}

func respondWith500(ctx *gin.Context, message string) {
	respondWithError(ctx, http.StatusInternalServerError, message)
}

func respondWithError(ctx *gin.Context, httpStatusCode int, message string) {
	err := Err{
		Error: message,
	}
	respondWithJSON(ctx, httpStatusCode, err)
}

func respondWithJSON(ctx *gin.Context, httpStatusCode int, jsonObject interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(httpStatusCode, jsonObject)
}
