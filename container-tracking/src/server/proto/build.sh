grpc_tools_node_protoc --plugin=protoc-gen-ts=../../../node_modules/.bin/protoc-gen-ts --ts_out=${PWD}/ server.proto \
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:${PWD}/ --grpc_out=${PWD}/ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` server.proto \
protoc --proto_path=../../../../api-gateway/pkg/tracking  --go_out=../../../../api-gateway/pkg/tracking  --go-grpc_out=../../../../api-gateway/pkg/tracking   -I ${PWD}/ server.proto \
