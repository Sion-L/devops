#!/bin/bash
goctl rpc protoc .\auth.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero