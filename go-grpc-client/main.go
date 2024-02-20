package main

import (
	"context"
	"encoding/json"
	"io"

	pb "go-grpc-client/proto/item"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		println("Failed to connect to GRPC server, cause: ", err.Error())
	}
	defer conn.Close()

	itemClient := pb.NewItemServiceClient(conn)
	getItemById(itemClient)
	getItemsByName(itemClient)
}

func getItemById(itemClient pb.ItemServiceClient) {
	response, _ := itemClient.GetItemById(context.Background(), &pb.ItemRequest{Id: 1})

	jsonByteArray, _ := json.MarshalIndent(response, "", "  ")

	println(string(jsonByteArray))
}

func getItemsByName(itemClient pb.ItemServiceClient) {
	stream, _ := itemClient.GetItemsByName(context.Background(), &pb.ItemRequest{Name: "sashimi"})
	for {
		item, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			println(err)
		}
		jsonByteArray, _ := json.MarshalIndent(item, "", "  ")

		println(string(jsonByteArray))

	}

}
