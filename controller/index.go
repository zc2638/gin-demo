package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct{ Base }

func (t *Index) Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

