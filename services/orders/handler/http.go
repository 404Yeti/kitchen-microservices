package handler

import (
	"html/template"
	"net/http"

	pb "kitchen.local/services/common/genproto/orders"
	"kitchen.local/services/orders/types"
	"kitchen.local/services/orders/util"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(ordersService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: ordersService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("/orders", h.CreateOrder)
	router.HandleFunc("/orders/view", h.ViewOrders)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req pb.CreateOrderRequest
	err := util.ParseJSONRequest(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &pb.Order{
		OrderID:    42,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &pb.CreateOrderResponse{
		Status: "Order created successfully",
	}
	util.WriteJSONResponse(w, http.StatusOK, res)
}


func (h *OrdersHttpHandler) ViewOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ordersService.GetOrders(r.Context(), 123) // replace with dynamic value if needed
	if err != nil {
		http.Error(w, "failed to fetch orders", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("orders").Parse(orderTemplate))
	tmpl.Execute(w, orders)
}


var orderTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Order List</title>
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
