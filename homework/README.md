# Mirco Service rebuild
-   BFF Service Admin Job Task
- API definition \ Error Code \ Error usage
- grpc
- Project structure 
- errgroup
- ELK Kafka Promethes
- consistent pipeline


Install grpc:
apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

Init mod:
go mod init github.com/haichaom/golang-geeke/homework
protoc api/log_process.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
