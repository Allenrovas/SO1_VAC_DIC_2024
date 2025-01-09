# GRPC

## 1. Instalacion de paquetes

```sh
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 2. Instalacion de protoc

```sh
sudo apt  install protobuf-compiler
```

## 3. Crear archivos compilados

```sh
protoc --go_out=. --go-grpc_out=. nombre_archivo.proto
```