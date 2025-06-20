package types

import (
    "context"
    pb "kitchen.local/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *pb.Order) error
    GetOrders(ctx context.Context, customerID int32) ([]*pb.Order, error)
}
