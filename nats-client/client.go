package main

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/SpeedyCoder/customers/utils"
)

const natsURI = "nats://0.0.0.0:4222"

func main() {
	nc, err := utils.GetNatsConnection(natsURI)
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	log.Println("Connected to NATS at:", nc.ConnectedUrl())

	var id uint64 = 1
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(id))

	start := time.Now()
	response, err := nc.Request("customers", data, 5*time.Second)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}
	duration := time.Since(start)
	log.Printf("Task scheduled in %+v Response: %v\n", duration, string(response.Data))
}
