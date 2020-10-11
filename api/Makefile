protoc: protoc-go protoc-js

protoc-go:
	mkdir -p pb/go
	protoc -I pb/v1/ \
		--go_out=plugins=grpc:pb/go \
		--gogrpcmock_out=:pb/go \
		pb/v1/*.proto

protoc-js:
	mkdir -p pb/js
	protoc -I pb/v1/ \
		--js_out=import_style=commonjs:pb/js \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:pb/js \
		pb/v1/*.proto

install:
	go get -u \
		github.com/golang/protobuf/proto \
		github.com/golang/protobuf/protoc-gen-go \
		google.golang.org/grpc \
		github.com/gogo/protobuf/protoc-gen-gogoslick \
		github.com/gogo/protobuf/gogoproto \
		github.com/onsi/ginkgo/ginkgo \
		github.com/onsi/gomega/...
	go install github.com/SafetyCulture/s12-proto/protobuf/protoc-gen-gogrpcmock

clean:
	rm ./pb/**/*.pb.go

lint:
	golangci-lint run

test:
	ginkgo -r -failFast

start-proxy:
	grpcwebproxy \
		--backend_addr=localhost:50051 \
		--run_tls_server=false \
		--allow_all_origins
