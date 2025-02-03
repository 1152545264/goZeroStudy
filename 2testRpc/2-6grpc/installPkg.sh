#protobuf和grpc核心代码下载
go get google.golang.org/protobuf@latest
go get google.golang.org/grpc@latest

#grpc的protoc代码生成工具
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#上面两个go install 安装完成之后,将${GOPATH}/bin/bin中的可执行文件复制到${GOPATH}/bin中,或者将${GOPATH}/bin/bin添加到系统环境变量

#protoc的编译器 protoc.exe NOTE:此处只是做一个示范,实际上此处应该下载的是Linux平台对应的二进制文件
wget https://github.com/protocolbuffers/protobuf/releases/download/v29.3/protoc-29.3-win64.zip
gunzip -k protoc-29.3-win64.zip
cd protoc-29.3-win64/bin && cp protoc.exe "${GOPATH}"/protoc.exe