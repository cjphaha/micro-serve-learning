package Weblib

import (
	"context"
	"fmt"
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
		prodRes,err := prodService.GetProdsList(context.Background(),&prodreq)
		c.JSON(200,gin.H{
			"data":prodRes,
			"err":err,
		})
	}
}