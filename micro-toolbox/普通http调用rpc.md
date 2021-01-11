可以设置micro网关
```shell
set MICRO_REHOSTRY=consul
set MICRO_REHOSTRY_ADDRESS=127.0.0.1:8500
set MICRO_API_NAMESPACE=api.jtthink.com
set MICRO_API_HANDLER=rpc
micro api
```
然后在postman里面就直接写成
post:    localhost:8080/test/TestService/call
tset是命名空间api.jtthink.com后面的TestService是服务名

然后当成正常的post用就可以了