package server

import (
	"fmt"

	pb "github.com/SpeedyCoder/customers/customers"
	"github.com/jinzhu/gorm"
)

// Customer model
type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
}

func (customer Customer) String() string {
	return fmt.Sprintf("|%s %s|", customer.FirstName, customer.LastName)
}

// Serialize Customer model to GRPC struct
func (customer *Customer) Serialize() *pb.Customer {
	return &pb.Customer{
		Id:        uint64(customer.ID),
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
	}
}
