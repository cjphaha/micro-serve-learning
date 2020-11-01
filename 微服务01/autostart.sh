#!/bin/sh
go run prod_main.go --server_address :8001&
go run prod_main.go --server_address :8002&
go run prod_main.go --server_address :8003
echo "执行成功"