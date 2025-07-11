package main
import (
    "log"
    "net"

    "google.golang.org/grpc"
	"kitchen.local/services/orders/handler"
	"kitchen.local/services/orders/service"
)

type gRPCServer struct {
	addr string 
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()

	//register our grpc services here
	orderService := service.NewOrdersService()
	handler.NewGrpcOrdersService(grpcServer, orderService)



	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}