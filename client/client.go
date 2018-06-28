package main

import (
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/SpeedyCoder/customers/customers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address   string = "localhost:5000"
	defaultID uint64 = 1
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCustomersClient(conn)

	// Contact the server and print out its response.
	id := defaultID
	if len(os.Args) > 1 {
		id, err = strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			log.Fatalf("%s is not a valid id", os.Args[1])
		}
	}
	log.Printf("Getting customer with id=%v", id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	customer, err := c.GetByID(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		log.Fatalf("could not find customer: %v", err)
	}
	log.Printf("Found customer with name \"%s %s\"", customer.GetFirstName(), customer.GetLastName())
}
