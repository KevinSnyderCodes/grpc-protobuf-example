#!/bin/bash

PROTOPATH=$GOPATH/src/github.com/kevinsnydercodes/protobuf-example/protobuf
FILES=$PROTOPATH/src/**/*.proto

for file in $FILES; do
  protoc \
    -I $PROTOPATH/src \
    --go_out="$PROTOPATH/compiled/go" \
    --ruby_out="$PROTOPATH/compiled/ruby" \
    --cpp_out="$PROTOPATH/compiled/cpp" \
    --php_out="$PROTOPATH/compiled/php" \
    --js_out="$PROTOPATH/compiled/js" \
    $file
done
