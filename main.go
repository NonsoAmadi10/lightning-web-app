package main

import (
	"context"
	"log"

	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10009", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to local lightning node: %v", err)
	}
	defer conn.Close()

	client := lnrpc.NewLightningClient(conn)
	ctx := context.Background()

	info, err := client.GetInfo(ctx, &lnrpc.GetInfoRequest{})
	if err != nil {
		log.Fatalf("failed to retrieve info from local lightning node: %v", err)
	}

	log.Println("Connected to local lightning node:", info)
}
