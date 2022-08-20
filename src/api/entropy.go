package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmillerv/gophrase/entropy"
	"net/http"
)

type submission struct {
	Phrase string `json:"passphrase"`
}

func postEntropy(c *gin.Context) {
	var submit submission
	err := c.BindJSON(&submit)
	if err != nil {
		c.JSON(http.StatusBadRequest, "improper json body")
	}
	metadata := entropy.GetEntropy(submit.Phrase)

	c.JSON(http.StatusOK, &metadata)
}
