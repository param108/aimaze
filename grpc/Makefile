SHELL := /bin/bash

go-protos: simulation.proto
	- mkdir -p go/maze
	protoc --go_out=./go/maze --go-grpc_out=. --go_opt=paths=source_relative simulation.proto

python-protos: simulation.proto python-env
	- mkdir -p python/maze
	source python-env/bin/activate; \
		python3 -m grpc_tools.protoc -I../grpc --python_out=./python/maze --grpc_python_out=./python/maze simulation.proto

python-env:
	python3 -m venv python-env
	source python-env/bin/activate; \
		pip3 install grpcio-tools grpcio

