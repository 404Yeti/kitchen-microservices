package main
import "google.golang.org/grpc"

func NewGRPCClient(addr string) *grpc.ClientConn  {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
}
    return conn, nil

func main() {

}
