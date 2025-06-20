package service

import (
	"context"
	pb "kitchen.local/services/common/genproto/orders"
)

var ordersDb = make([]*pb.Order, 0)

type OrdersService struct {
	// can add fields like a DB connection later
}

func NewOrdersService() *OrdersService {
	return &OrdersService{}
}

func (s *OrdersService) CreateOrder(ctx context.Context, order *pb.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrdersService) GetOrders(ctx context.Context, customerID int32) ([]*pb.Order, error) {
	var result []*pb.Order
	for _, o := range ordersDb {
		if o.CustomerID == customerID {
			result = append(result, o)
		}
	}
	return result, nil
}