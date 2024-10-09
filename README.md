# gomicroBDD


raíz del proyecto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



raíz del proyecto
protoc --go_out=. --go-grpc_out=. api/proto/person_account.proto


go mod init gomicrobase
go mod tidy

//*Para actualizar Dependencias*/
go build
go get -u

//Inicializar el servicio
go run cmd/server/main.go


//Abrir una nueva terminal apra probar el servicio
cd .\test\
go run .\servicetest.go




