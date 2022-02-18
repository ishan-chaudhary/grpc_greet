# Use this command to generate grpc and go files from protobuf files
protoc greet/greetpb/greet.proto --go-grpc_out=. --go_out=.

#Server Start
go run greet/greet_server/server.go

#Client Start
go run greet/greet_client/client.go

#Protobuf for calculator
protoc calculator/calculatorpb/calculator.proto --go-grpc_out=. --go_out=.
go run calculator/calculator_client/client.go
go run calculator/calculator_server/server.go

#Protobuf for blog
protoc blog/blogpb/blog.proto --go-grpc_out=. --go_out=.
go run calculator/calculator_client/client.go
go run calculator/calculator_server/server.go