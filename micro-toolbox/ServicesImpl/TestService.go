package ServicesImpl

import (
	"context"
	"micro-toolbox/Services"
	"strconv"
)

type TestService struct{

}
//call接口在test.pb.micro.go离main
func(this *TestService)Call(ctx context.Context, req *Services.TestRequest, rsp *Services.TestResponse) error{
	rsp.Data = "test" + strconv.Itoa(int(req.Id))
	return nil
}