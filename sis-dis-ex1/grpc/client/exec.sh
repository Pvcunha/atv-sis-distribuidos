#!/bin/bash

# go run ../server/server.go &
sleep 5
echo "step 1"
go run client.go -run=100 -conc=true > grpc_conc_100.txt
go run client.go -run=100 -conc=false > grpc_100.txt
echo "step 2"
go run client.go -run=1000 -conc=true > grpc_conc_1000.txt
go run client.go -run=1000 -conc=false > grpc_1000.txt
echo "step 3"
go run client.go -run=10000 -conc=true > grpc_conc_10000.txt
go run client.go -run=10000 -conc=false > grpc_10000.txt
