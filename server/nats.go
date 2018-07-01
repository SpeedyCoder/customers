package server

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/SpeedyCoder/customers/utils"
	"github.com/jinzhu/gorm"
	"github.com/nats-io/go-nats"
)

const natsURI = "nats://0.0.0.0:4222"

func serveCustomer(db *gorm.DB, nc *nats.Conn) nats.MsgHandler {
	return func(message *nats.Msg) {
		var customer Customer
		id := binary.BigEndian.Uint64(message.Data)
		dbc := db.First(&customer, id)
		if dbc.Error != nil {
			log.Printf("NatsByID id=%v NOT FOUND", id)
			nc.Publish(message.Reply, []byte("Error|NotFound"))
			return
		}
		log.Printf("NatsByID id=%v OK", id)
		response := fmt.Sprintf(
			"Customer|%v|%s|%s",
			customer.ID,
			customer.FirstName,
			customer.LastName,
		)
		nc.Publish(message.Reply, []byte(response))
	}
}

// ServeNats serves requests from NATS.
func ServeNats(db *gorm.DB) {
	nc, err := utils.GetNatsConnection(natsURI)
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	fmt.Println("Connected to NATS at:", nc.ConnectedUrl())
	nc.Subscribe("customers", serveCustomer(db, nc))

	fmt.Println("Subscribed to 'customers' for processing requests...")
}
