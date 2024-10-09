package main

import (
	"context"
	api "gomicrobase/api/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewPersonAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	persona := &api.Person{Name: "Juan", Email: "miaCj@example.com", Id: 1}
	personResponse, err := c.GetPerson(ctx, persona)
	if err != nil {
		log.Fatalf("could not get person: %v", err)
	}
	log.Printf("Name: %s, Email: %s", personResponse.Name, personResponse.Email)
}
