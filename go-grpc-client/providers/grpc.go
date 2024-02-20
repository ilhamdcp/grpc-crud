package providers

import (
	"go-grpc-client/proto/item"
	pb "go-grpc-client/proto/item"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	ItemClient item.ItemServiceClient
}

func ProvideClientConnection() GrpcClient {
	conn, err := grpc.Dial("http://localhost:8090")
	if err != nil {
		println("Failed to connect to GRPC server, cause: ", err.Error())
	}
	defer conn.Close()

	return GrpcClient{
		ItemClient: pb.NewItemServiceClient(conn),
	}
}
