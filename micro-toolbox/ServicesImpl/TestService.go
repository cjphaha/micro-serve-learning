package ServicesImpl
/*
实现类，实现services/protos
 */
import (
	"context"
	"micro-toolbox/Services"
	"strconv"
)

type TestService struct{

}
//call接口在test.pb.micro.go里面，这里是具体的实现
func(this *TestService)Call(ctx context.Context, req *Services.TestRequest, rsp *Services.TestResponse) error{
	rsp.Data = "test" + strconv.Itoa(int(req.Id))
	return nil
}