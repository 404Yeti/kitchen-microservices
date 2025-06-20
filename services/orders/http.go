package main

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "kitchen.local/services/common/genproto/orders"
	"google.golang.org/grpc"
)

type httpServer struct {
	addr string
}

func NewHTTPServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := pb.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		_, err := c.CreateOrder(ctx, &pb.CreateOrderRequest{
			CustomerID: 123,
			ProductID:  456,
			Quantity:   2,
		})
		if err != nil {
			log.Printf("client error: %v", err)
			http.Error(w, "failed to create order", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Order created via gRPC!"))
	})

	log.Println("Starting HTTP server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

// You can later move this to a separate file if needed
var orderTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Order Form</title>
</head>
<body>
	<h1>Order List</h1>
	<table border="1">
		<tr>
			<th>Order ID</th>
			<th>Customer ID</th>
			<th>Product ID</th>
			<th>Quantity</th>
		</tr>
		{{range .}}
		<tr>
			<td>{{.OrderID}}</td>
			<td>{{.CustomerID}}</td>
			<td>{{.ProductID}}</td>
			<td>{{.Quantity}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>`
