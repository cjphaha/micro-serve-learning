package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"micro-toolbox/logs"
	"micro-toolbox/sidecar"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.POST("/testpost", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"mes":"ok",
			})
		})
	}
	server := &http.Server{
		Addr: ":8088",
		Handler: ginRouter,
	}

	service := sidecar.NewService("com.cjp.gin2")
	service.AddNode("test-" + uuid.New().String(),8088,"localhost")

	handler := make(chan error)
	go(func(){
		handler <- server.ListenAndServe()//如果有错误把err给handler
	})()
	go(func() {
		//创建信号监听
		notify := make(chan os.Signal)
		signal.Notify(notify,syscall.SIGTERM,syscall.SIGINT,syscall.SIGKILL)
		handler <- fmt.Errorf("%s",<-notify)//监听到有值，就把值给handler
	})()
	go(func() {
		//注册服务
		err := sidecar.RegService(service)
		if err != nil{
			handler <- err
		}
	})()


	getHandler := <- handler
	logs.Log.Println(getHandler)

	//反注册
	err := sidecar.UnRegService(service)
	if err != nil {
		logs.Log.Println(err)
	}
	err = server.Shutdown(context.Background())
	if err != nil{
		logs.Log.Println(err)
	}
}
