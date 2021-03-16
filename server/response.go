package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int `json:"message"`
}

func accepted(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{Status: 0})
}

func badRequest(c *gin.Context, err string) {
	log.Warn(err)
	c.JSON(http.StatusBadRequest, &Response{Status: 1})
}

func internalError(c *gin.Context, err string) {
	log.Error(err)
	c.JSON(http.StatusInternalServerError, &Response{Status: 2})
}
