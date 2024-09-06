#!/bin/bash
goctl rpc protoc ./user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero

