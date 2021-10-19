package service

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Meta map[string]interface{} `json:"meta"`
	Data interface{}            `json:"data"`
}

func Response(data interface{}, c *gin.Context, table string, total int64) interface{} {
	var metaData map[string]interface{}

	result := response{
		Meta: metaData,
		Data: data,
	}

	return result
}
