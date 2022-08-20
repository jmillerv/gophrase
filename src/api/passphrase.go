package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmillerv/gophrase/corpus"
	"github.com/jmillerv/gophrase/entropy"
	"github.com/jmillerv/gophrase/generate"
	"net/http"
	"strconv"
)

type response struct {
	Passphrase string            `json:"passphrase"`
	Metadata   *entropy.Metadata `json:"metadata"`
}

func getPassphrase(c *gin.Context) {
	params := httpParamsToStruct(c)
	phrase := generate.Password(params)
	metadata := entropy.GetEntropy(phrase)
	c.JSON(http.StatusOK, &response{
		Passphrase: phrase,
		Metadata:   metadata,
	})
}

func httpParamsToStruct(c *gin.Context) *generate.Params {
	count, _ := strconv.Atoi(c.Query("word_count"))
	capitals, _ := strconv.ParseBool(c.Query("capital"))
	special, _ := strconv.ParseBool(c.Query("special"))
	number, _ := strconv.ParseBool(c.Query("number"))
	params := &generate.Params{
		WordCount:    count,
		WordList:     corpus.SetWordList(c.Query("word_list")),
		Capitals:     capitals,
		SpecialChars: special,
		Numbers:      number,
	}
	return params
}
