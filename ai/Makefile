SHELL := /bin/bash

setup:
	source env/bin/activate; \
		pip3 install -r /path/to/requirements.txt

run:
	PYTHONPATH="python/maze" python3 simulate.py

python-protos: ../grpc/simulation.proto env
	- mkdir -p python/maze
	source env/bin/activate; \
		python3 -m grpc_tools.protoc -I../grpc --python_out=./python/maze --grpc_python_out=./python/maze simulation.proto

save:
	pip3 freeze > requirements.txt


