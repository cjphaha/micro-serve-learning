package Weblib

import (
	"github.com/gin-gonic/gin"
	"miniserver/Services"
)

func NewGinRouter(prodService Services.ProdService) *gin.Engine{
	ginRouter := gin.Default()
	ginRouter.Use(Initmiddleware(prodService),ErrorMiddleware())
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods",GetProdsList)
		v1Group.Handle("GET", "/prods/:ProdId",GetProdDetail)
	}
	return ginRouter
}
