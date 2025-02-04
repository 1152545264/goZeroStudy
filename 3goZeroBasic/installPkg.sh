go install github.com/zeromicro/go-zero/tools/goctl@latest
go get github.com/zeromicro/go-zero@latest

#安装etcd和redis，下载解压后将目录添加到环境变量中
# etcd下载地址: https://github.com/etcd-io/etcd/releases/
#windows版本redis下载地址：https://github.com/redis-windows/redis-windows/releases
#linux上直接apt或者yum安装即可

#使用goctl创建user微服务：创建user目录之后，执行下面的一条命令
goctl rpc new user

#根据指定proto创建微服务项目
goctl rpc  protoc .\user.proto --go_out=. --go-grpc_out=. --zrpc_out=.