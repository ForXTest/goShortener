package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Status int32 `json: status`
	Error string `json:"error"`
}

func responseBadRequest(context *gin.Context, message string) {
	context.JSON(
		http.StatusBadRequest,
		errorResponse{http.StatusBadRequest, message},
	)
}

func responseServerError(context *gin.Context) {
	context.JSON(
		http.StatusInternalServerError,
		errorResponse{http.StatusInternalServerError, "Internal server error"},
	)
}

func responseNotFound(context *gin.Context) {
	context.JSON(
		http.StatusNotFound,
		errorResponse{http.StatusNotFound, "Not Found"},
	)
}
