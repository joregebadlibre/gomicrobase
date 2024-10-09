package main

import (
	"log"
	"net"

	api "gomicrobase/api/proto"
	"gomicrobase/config"
	"gomicrobase/internal/db"
	"gomicrobase/internal/handler"
	"gomicrobase/internal/repository"
	"gomicrobase/internal/service"

	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	dbConn, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}
	log.Default().Println("connect to database: ", dbConn)
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	repo := repository.NewPersonAccountRepository(dbConn)
	personAccountService := service.NewPersonAccountService(repo)
	personAccountHandler := handler.NewPersonAccountHandler(personAccountService)

	// Registra el servicio en el servidor gRPC
	api.RegisterPersonAccountServiceServer(s, personAccountHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
