package Weblib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniserver/Services"
)
//添加统一处理错误的中间件
func Initmiddleware(prodService Services.ProdService) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodService"] = prodService//赋值
		c.Next()
	}
}

func PanicIfError(err error){
	if err!=nil{
		panic(err)
	}
}

func ErrorMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if r := recover();r!=nil{
				c.JSON(500,gin.H{
					"status":fmt.Sprintf("%s",r),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
