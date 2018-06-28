package server

import (
	"context"
	"errors"
	"log"

	pb "github.com/SpeedyCoder/customers/customers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // We use postgress
)

// Server struct
type Server struct {
	DB *gorm.DB
}

// GetByID GRPC method
func (s *Server) GetByID(context context.Context, request *pb.IdRequest) (*pb.Customer, error) {
	var customer Customer
	dbc := s.DB.First(&customer, request.Id)
	log.Println(dbc)
	if dbc.Error != nil {
		log.Printf("GetByID id=%v NOT FOUND", request.Id)
		return nil, errors.New("Not found")
	}
	log.Printf("GetByID id=%v OK", request.Id)
	return customer.Serialize(), nil
}

// List GRPC method
func (s *Server) List(request *pb.ListRequest, listServer pb.Customers_ListServer) error {
	log.Printf("List limit=%v offset=%v\n", request.Limit, request.Offset)

	var customers []Customer
	s.DB.Offset(request.Offset).Limit(request.Limit).Find(&customers)

	for _, customer := range customers {
		err := listServer.Send(customer.Serialize())
		if err != nil {
			return err
		}
	}
	return nil
}
