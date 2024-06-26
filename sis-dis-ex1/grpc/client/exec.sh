#!/bin/bash

go run ../server/server.go &
sleep 5
echo "step 1"
go run client.go -run=100 > grpc_100.txt
echo "step 2"
go run client.go -run=1000 > grpc_1000.txt
echo "step 3"
go run client.go -run=10000 > grpc_10000.txt