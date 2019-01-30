all: proto-go proto-py3

proto-go:
	protoc -I ./proto --go_out=plugins=grpc:. ./proto/opensnitch/network/*.proto
	protoc -I ./proto --go_out=plugins=grpc:. ./proto/opensnitch/rules/*.proto
	protoc -I ./proto --go_out=plugins=grpc:. ./proto/opensnitch/ui/*.proto
	cp -r github.com/evilsocket/opensnitch/* .
	rm -rf github.com/

proto-py3:
	mkdir -p python3
	python3 -m grpc_tools.protoc -I ./proto \
		--python_out=./python3/ \
		./proto/opensnitch/network/*.proto
	python3 -m grpc_tools.protoc -I ./proto \
		--python_out=./python3/ \
		./proto/opensnitch/rules/*.proto
	python3 -m grpc_tools.protoc -I ./proto \
		--python_out=./python3/ \
		--grpc_python_out=./python3/ \
		./proto/opensnitch/ui/*.proto
	find ./python3/opensnitch -type d -exec touch {}/__init__.py \;