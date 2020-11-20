package sidecar

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"micro-toolbox/logs"
	"net/http"
)

/*
json rpc
 */
type JSONRequest struct {
	Jsonspc string
	Method string
	Params []*Service
	Id i
func NewJSONRequest(service *Service,endpoinr string) *JSONRequest{
	return &JSONRequest{Jsonspc: "2.0",Method: endpoinr,Params: []*Service{service},Id: 1}
}

var (
	RegistryURL = "http://localhost:2379"
)
/*
注册的主体方法
 */

func requestRegistry(jsonrequest *JSONRequest) error {
	b,err := json.Marshal(jsonrequest)//jsonrequest变成json字符串
	if err!= nil{
		logs.Log.Println(err)
		return err
	}
	rsp,err := http.Post(RegistryURL,"application/json",bytes.NewBuffer(b))
	if err != nil{
		logs.Log.Println(err)
	}
	defer rsp.Body.Close()
	res,err := ioutil.ReadAll(rsp.Body)//打印结果
	if err != nil{
		return err
	}
	logs.Log.Println(string(res) + "\n" + string(b))
	return nil
}
//删除
func UnRegService(service *Service) error{
	return requestRegistry(NewJSONRequest(service,"Registry.Deregister"))
}
//注册
func RegService(service *Service) error{
	return requestRegistry(NewJSONRequest(service,"Registry.Register"))
}