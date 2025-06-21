# Kitchen Microservices — Go + gRPC + HTTP

A lightweight microservices application using **Go**, **gRPC**, and **HTML templating** — built to manage and view customer orders.

---

## Tech Stack

-  Go 1.22+
-  gRPC + Protocol Buffers
-  HTML rendering with `html/template`
-  In-memory order storage
-  Makefile-based protobuf generation

---

## 📁 Project Structure
kitchen/
├── protobuf/ # .proto definitions
├── services/
│ ├── orders/
│ │ ├── handler/ # gRPC + HTTP handlers
│ │ ├── service/ # Business logic layer
│ │ ├── types/ # Interface definitions
│ │ ├── util/ # Helpers
│ │ └── main.go # Entrypoint
├── Makefile
└── README.md

--

## Usage

### Generate Protobuf Code
make gen

## Run the App
make run-orders

## Sample EndPoints
create a new order
POST /Orders
curl -X POST http://localhost:8000/orders \
  -H "Content-Type: application/json" \
  -d '{"customerID":1,"productID":2,"quantity":3}'

  GET /orders/View
  http://localhost:8000/orders/view?customerID=1