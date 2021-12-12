FROM ubuntu20.04
USER root
RUN mkdir -p /usr/local/gopath
RUN wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz -O /tmp/go1.17.5.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go1.17.5.linux-amd64.tar.gz
RUN source envionment
apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
protoc api/log_process/log_process.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
go build cmd/server/main.go
ENTRYPOINT ./main
