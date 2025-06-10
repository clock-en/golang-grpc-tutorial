package main

import (
	"context"
	"fmt"
	"github.com/clock-en/golang-grpc-tutorial/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// NOTE: 非 SSL になるので本番では使用できない
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	callListFiles(client)
}

func callListFiles(client pb.FileServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFilenames())
}
