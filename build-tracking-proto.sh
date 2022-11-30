protoc --proto_path=${PWD}/protobuf/tracking/  --go_out=${PWD}/protobuf/tracking/ --go-grpc_out=${PWD}/protobuf/tracking/ --js_out="import_style=commonjs,binary:${PWD}/container-tracking/src/server/proto" --ts_out="${PWD}/container-tracking/src/server/proto" -I ${PWD}/protobuf/tracking tracking.proto
cd container-tracking
npm run proto-gen