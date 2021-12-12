# Mirco Service rebuild

Tech points:
-   BFF Service Admin Job Task
- API definition \ Error Code \ Error usage
- grpc 
- Project structure 
- errgroup
- ELK Kafka Promethes
- consistent pipeline

# Project description

## Log processing based grpc.

Get useful error message from the log file. Log file is in a nfs server, and it is
not allowed mount from client, so we need get call grpc service in the server.

Currently, it is python based, rebuild to modern go app.


##Install grpc:
apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
protoc api/log_process/log_process.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

##Init mod:
go mod init github.com/haichaom/golang-geeke/homework


## Run
Server:
. environment
go run cmd/server/main.go


Client:
. environment
go run cmd/client/main.go
