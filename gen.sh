#!/bin/bash

protoc  -I/usr/local/include -I. \
  -I $GOPATH/src \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --go_opt=paths=source_relative \
  --grpc-gateway_out=logtostderr=true,repeated_path_param_separator=ssv:. \
  --swagger_out=logtostderr=true,repeated_path_param_separator=ssv:. \
  qrcode/cmd/grpcapps/pob/*.proto