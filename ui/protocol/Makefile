all: ui.pb.go

ui.pb.go:
	protoc -I. ui.proto --go_out=plugins=grpc:.

clean:
	@rm -rf ./*.pb.go