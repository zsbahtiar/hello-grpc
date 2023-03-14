# hello-grpc

## gen protobuf
```shell
protoc \
    --proto_path=../proto \
    --go_out=plugins=grpc:./handler/api/grpc/{service} \
    --grpc-gateway_out=logtostderr=true:../handler/api/grpc/{service} \
    ../proto/{x}.proto
```