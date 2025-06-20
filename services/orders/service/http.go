package service

import (
	"log"
	"net/http"

	"kitchen.local/services/orders/handler"
)

type httpServer struct {
	addr string
}

func NewHTTPServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := NewOrdersService() // ✅ use directly
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting HTTP server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
