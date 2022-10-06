

protos: simulation.proto
	- rm -rf go
	mkdir -p go/maze
	protoc --go_out=./go/maze --go-grpc_out=. --go_opt=paths=source_relative simulation.proto
