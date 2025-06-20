package handler

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "kitchen.local/services/common/genproto/orders"
	"kitchen.local/services/orders/types"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	pb.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpcServer *grpc.Server, ordersService types.OrderService) *OrdersGrpcHandler {
	handler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return handler
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &pb.Order{
		OrderID:    42,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	if err := h.ordersService.CreateOrder(ctx, order); err != nil {
		log.Printf("failed to create order: %v", err)
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Status: "Order created successfully",
	}, nil
}


func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrderResponse, error) {
	orders, err := h.ordersService.GetOrders(ctx, req.CustomerID)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{
		Orders: orders,
	}, nil
}
