package sidecar
/*
主要是为了注册方便
 */
type Service struct {
	Name string
	Nodes []*ServiceNode
}

type ServiceNode struct {
	Id string //服务id不能重复
	Port int
	Address string
}

//实例化service
func NewService(name string) *Service{
	return &Service{Name: name,Nodes: make([]*ServiceNode,0)}
}

func NewServiceNode(id string,port int,address string) *ServiceNode{
	return &ServiceNode{Id: id,Port: port,Address: address}
}

func(this *Service) AddNode(id string,port int,address string){
	this.Nodes = append(this.Nodes,NewServiceNode(id,port,address))
}