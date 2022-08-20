package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmillerv/gophrase/corpus"
	"net/http"
)

func getOptions(c *gin.Context) {
	c.String(http.StatusOK, string(corpus.PrintWordListOptions()))
}
