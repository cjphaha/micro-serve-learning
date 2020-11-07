package Weblib

import (
	"github.com/gin-gonic/gin"
	"miniserver/Services"
)

func Initmiddleware(prodService Services.ProdService) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodService"] = prodService//赋值
		c.Next()
	}
}
