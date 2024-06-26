#!/bin/bash
go run ../server/server.go &
sleep 5
echo "step 1"
go run client.go -run=100 > gorpc_100.txt
echo "step 2"
go run client.go -run=1000 > gorpc_1000.txt
echo "step 3"
go run client.go -run=10000 > gorpc_10000.txt