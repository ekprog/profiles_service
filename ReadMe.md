## About

This project is example of Profile Service using clean design for gRPC microservices.
Template for clean design you can find [HERE](https://github.com/ekprog/microservice_clean_design)

**We have 4 main layer:**
1) Deliver layer
2) Interactors (Use Cases)
3) Repositories
4) Domain

**1** -> **2** -> **3** -> **4**

*First layer can use any other, Second layer can use only layers 3-4 and so on.*


### Domain

Models, UseCase behaviour (interfaces). 
!! NO IMPLEMENTATION !!

### Repositories

Layer for getting data from any source. In example you can see postgres DB and test Mock implementation.

### Interactors

Here we implement main logic of our service. Whole implementation depends on domain layer behaviour.

### Delivery layer

After running of UseCases we should send result back to client. Here we can implement different clients types. \
For example, gRPC client, REST client, CLI client. \
So this layer doesn't depend on implementation of our core (interactors) and easy for testing.

## Examples

Full example with JWT Auth microservice using this clean design you can find here - [https://github.com/ekprog/grpc_auth_service](https://github.com/ekprog/grpc_auth_service)

## Compile gRPC files

```go
protoc -I ./proto \
--go_out ./pkg/pb \
--go_opt paths=source_relative \
--go-grpc_out ./pkg/pb \
--go-grpc_opt paths=source_relative \
--grpc-gateway_out ./pkg/pb \
--grpc-gateway_opt paths=source_relative \
--openapiv2_out ./docs \
./proto/api/**/*.proto

```