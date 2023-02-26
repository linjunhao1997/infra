# 第一个I为google的proto包路径，第二个I为自己引用的proto包路径
goctl rpc protoc -I /d:/learning  -I ../desc ../desc/system.proto --go_out=../pb --go-grpc_out=../pb  --zrpc_out=../