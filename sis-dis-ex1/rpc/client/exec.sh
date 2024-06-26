#!/bin/bash
go run ../server/server.go &

go run client.go -run=100 > gorpc_100.txt
go run client.go -run=1000 > gorpc_1000.txt
go run client.go -run=10000 > gorpc_10000.txt