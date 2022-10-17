SHELL := /bin/bash

go-protos: ../grpc/simulation.proto
	- mkdir -p spec/grpc/maze
	protoc -I ../grpc --go_out=./spec/grpc/maze --go-grpc_out=. --go_opt=paths=source_relative ../grpc/simulation.proto

