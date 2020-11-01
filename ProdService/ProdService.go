package ProdService

import "strconv"

type ProdModel struct {
	ProdID   int
	ProdName string
}

func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{ProdName: pname, ProdID: id}
}

func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "ProdName"+strconv.Itoa(i)))
	}
	return ret
}
