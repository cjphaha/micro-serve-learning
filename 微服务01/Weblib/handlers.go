package Weblib

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"miniserver/Services"
)

func GetProdsList(c *gin.Context)  {
	prodService := c.Keys["prodService"].(Services.ProdService)//后面这个(Services.ProdService)是类型断言
	//因为编译器不知道c.Keys["prodService"]取出来的是什么类型，如果不加类型的话21行就会无法执行GetProdsList,所以要加一个类型
	//断言，编译器认为这个是Services.ProdService类型，学到了学到了
	var prodreq Services.ProdsRequest
	err := c.Bind(&prodreq)
	fmt.Println(prodreq)
	if err != nil{//如果出错
		c.JSON(500,gin.H{
			"status":err.Error(),
		})
	}else {
		//熔断代码改造
		//设置config
		configA := hystrix.CommandConfig{
			Timeout: 1000,
		}
		var prodRes *Services.ProdListResponse
		//配置command
		hystrix.ConfigureCommand("getprods",configA)
		//执行，实行do方法
		err := hystrix.Do("getprods", func() error {
			prodRes,err = prodService.GetProdsList(context.Background(),&prodreq)
			return err
		},nil)
		if err != nil{
			c.JSON(500,gin.H{
				"status":err.Error(),
			})
		}
		c.JSON(200,gin.H{
			"data":prodRes,
			"err":err,
		})
	}
}