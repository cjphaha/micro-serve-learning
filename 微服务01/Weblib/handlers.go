package Weblib

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"miniserver/Services"
	"strconv"
)

//测试方法
func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdName: pname, ProdID: id}
}

func defaultProds() (*Services.ProdListResponse,error){
	models := make([]*Services.ProdModel,0)
	var i int32
	for i = 0;i < 5;i ++{
		models = append(models,newProd(100 + i,"prodname" + strconv.Itoa(100 + int(i))))
	}
	res := &Services.ProdListResponse{}
	res.Data = models
	return res,nil
}

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
		if err != nil{
			c.JSON(500,gin.H{
				"status":err.Error(),
			})
		}
		c.JSON(200,gin.H{
			"data":prodRes,
			"err":err,
		})
		////熔断代码改造
		////设置config
		//configA := hystrix.CommandConfig{
		//	Timeout: 1000,
		//}
		//var prodRes *Services.ProdListResponse
		////配置command
		//hystrix.ConfigureCommand("getprods",configA)
		////执行，实行do方法
		//err := hystrix.Do("getprods", func() error {
		//	prodRes,err = prodService.GetProdsList(context.Background(),&prodreq)
		//	return err
		//}, func(e error) error {
		//	prodRes,err = defaultProds()
		//	return err
		//})//fallback降级处理
		//if err != nil{
		//	c.JSON(500,gin.H{
		//		"status":err.Error(),
		//	})
		//}
		//c.JSON(200,gin.H{
		//	"data":prodRes,
		//	"err":err,
		//})
	}
}

func GetProdDetail(c *gin.Context){
	var prodReq Services.ProdsRequest
	//binduri是针对于gei请求的
	PanicIfError(c.BindUri(&prodReq))
	prodService := c.Keys["prodService"].(Services.ProdService)
	resp,_ := prodService.GetProdsDetail(context.Background(),&prodReq)
	c.JSON(200,gin.H{
		"data":resp.Data,
	})
}