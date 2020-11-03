package ServiceImpl

import (
	"context"
	"micro-serve-2/Services"
)

type newProd(id int32,pname string) *Services.ProdModel{
	return &Services.ProdModel ()
}

type ProdService struct {
	
}

func(*ProdService) GetProdsList(ctx context.Context,in *Services.ProdsRequest,res *Services.ProdListResponse) error{

}