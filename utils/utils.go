package utils

import (
	"fmt"
	"time"

	nats "github.com/nats-io/go-nats"
)

// GetNatsConnection returns a conection to nats
func GetNatsConnection(natsURI string) (*nats.Conn, error) {
	var err error
	var nc *nats.Conn

	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(natsURI)
		if err == nil {
			return nc, nil
		}

		fmt.Println("Waiting before connecting to NATS at:", natsURI)
		time.Sleep(1 * time.Second)
	}
	return nil, err
}
