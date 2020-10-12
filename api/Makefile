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

install-grpcwebproxy:
	GOPATH=~/go ; export GOPATH
	git clone https://github.com/improbable-eng/grpc-web.git $GOPATH/src/github.com/improbable-eng/grpc-web
	cd $GOPATH/src/github.com/improbable-eng/grpc-web
	dep ensure # after installing dep
	go install ./go/grpcwebproxy # installs into $GOPATH/bin/grpcwebproxyfrd

install-grpc-cli:
	git clone https://github.com/grpc/grpc
	sudo apt-get install libgflags-dev
	cd grpc
	git submodule update --init
	mkdir -p cmake/build
	cd cmake/build
	cmake -DgRPC_BUILD_TESTS=ON ../..
	make grpc_cli

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

db.migrate.up:
	goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:4432/plant-manager?connect_timeout=180&sslmode=disable" up
	goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:4432/plant-manager-test?connect_timeout=180&sslmode=disable" up

db.migrate.down:
	goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:4432/plant-manager?connect_timeout=180&sslmode=disable" reset
	goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:4432/plant-manager-test?connect_timeout=180&sslmode=disable" reset

