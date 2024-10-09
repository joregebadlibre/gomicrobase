package main

import (
	"log"
	"net"

	"gomicrobase/config"
	"gomicrobase/internal/db"
	"gomicrobase/internal/handler"

	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	dbConn, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	handler.RegisterPersonAccountServiceServer(grpcServer, handler.NewPersonAccountHandler(dbConn))

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
