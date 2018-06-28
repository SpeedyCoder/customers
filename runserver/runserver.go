package main

import (
	"log"
	"net"

	pb "github.com/SpeedyCoder/customers/customers"
	"github.com/SpeedyCoder/customers/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":5000"

func main() {
	db := server.GetDB()
	server.MigrateDB(db)
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomersServer(s, &server.Server{db})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
