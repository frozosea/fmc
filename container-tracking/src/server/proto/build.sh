grpc_tools_node_protoc --plugin=protoc-gen-ts=../../../node_modules/.bin/protoc-gen-ts --ts_out=${PWD}/ --grpc_out=${PWD}/ --go-grpc_out=${PWD} --js_out=import_style=commonjs,binary:${PWD}/ tracking.proto \
##grpc_tools_node_protoc --js_out=import_style=commonjs,binary:${PWD}/ --grpc_out=${PWD}/  tracking.proto \
#protoc --proto_path=${PWD}  --go_out=${PWD}/ --go-grpc_out=${PWD}/ --js_out="import_style=commonjs,binary:${PWD}/" --ts_out="${PWD}/" -I ${PWD}/ tracking.proto

PROTOC_GEN_TS_PATH="../../../node_modules/.bin/protoc-gen-ts"
#
## Directory to write generated code to (.js and .d.ts files)
##OUT_DIR="./generated"
#
#protoc \
#    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
#    --js_out="import_style=commonjs,binary:${PWD}/" \
#    --ts_out="${OUT_DIR}" \
#    -I ${PWD}/ tracking.proto
