package main

/*
测试函数
*/

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	microhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"miniserver/model"
	"net/http"
)

//比较推荐的调用方式
func callAPI2(s selector.Selector) {
	myclient := microhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"), //m哦人的prutof
	)
	req := myclient.NewRequest("myConsul", "/v1/prods",
		model.ProdsRequest{
			Size: 3,
		})
	var rsp model.ProdListResponse
	err := myclient.Call(context.Background(), req, &rsp) //rsp是返回的对象,传地址的话可以解析道rsp
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.GetData())
}

//原始调用方式
func callAPI(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	callAPI2(mySelector)
}
