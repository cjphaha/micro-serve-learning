package ServiceImpl
//是一个实现类
//对prodservice.micro.go 中service的具体的实现

import (
	"context"
	"micro-serve-2/Services"
	"strconv"
	"time"
)

//测试方法
func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdName: pname, ProdID: id}
}

type ProdService struct {

}

func(*ProdService) GetProdsList(ctx context.Context,in *Services.ProdsRequest,res *Services.ProdListResponse) error{
	time.Sleep(time.Second*3)
	models := make([]*Services.ProdModel,0)
	var i int32
	for i = 0;i < in.Size;i ++{
		models = append(models,newProd(100 + i,"prodname" + strconv.Itoa(100 + int(i))))
	}
	res.Data = models
	return nil
}


func(*ProdService) GetProdsDetail(ctx context.Context, req *Services.ProdsRequest, rsp *Services.ProdDetailResponse) error {
	rsp.Data = newProd(req.ProdId,"商品详情")
	return nil
}
