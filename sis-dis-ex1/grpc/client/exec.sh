#!/bin/bash

go run ../server/server.go &

go run client.go -run=100 > grpc_100.txt
go run client.go -run=1000 > grpc_1000.txt
go run client.go -run=10000 > grpc_10000.txt