.PHONY: requirements api-external

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

requirements:
	go mod download
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/golang/protobuf/protoc-gen-go

api-external:
	# gRPC
	protoc -I=$(GOOGLEAPIS_PATH) -I=../../protobuf/external --go_out=plugins=grpc,paths=source_relative:. ../../protobuf/external/device.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../../protobuf/external --go_out=plugins=grpc,paths=source_relative:. ../../protobuf/external/measurement.proto

	protoc -I=$(GOOGLEAPIS_PATH) -I=../../protobuf/external --grpc-gateway_out=paths=source_relative,logtostderr=true:. ../../protobuf/external/device.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../../protobuf/external --grpc-gateway_out=paths=source_relative,logtostderr=true:. ../../protobuf/external/measurement.proto
