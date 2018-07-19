protoc -I. \
-I${GOPATH}/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:${GOPATH}/src \
--validate_out="lang=go:${GOPATH}/src" \
demo/grpc/user_service.proto \

protoc -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
grpc/user_service.proto
