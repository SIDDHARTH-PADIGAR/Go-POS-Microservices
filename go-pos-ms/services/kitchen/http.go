package main

import (
	"context"
	"go-pos-ms/services/common/genproto/orders"
	"log"
	"net/http"
	"text/template"
	"time"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":8080")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Added 'r *http.Request'
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3123,
			Quantity:   2,
		})

		if err != nil { // Fixed 'nill' to 'nil'
			log.Fatal("Client error: %v", err)
		}

		res, err := c.GetOrders(ctx, &orders.GetOrderRequest{
			CustomerID: 42,
		})
		if err != nil {
			log.Fatal("Client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatal("Template error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>POS Orders</title>
    <style>
        body { font-family: Arial, sans-serif; background: #121212; color: #e0e0e0; padding: 20px; }
        h1 { text-align: center; color: #ffcc00; }
        table { width: 100%; border-collapse: collapse; }
        th, td { padding: 10px; border: 1px solid #444; text-align: center; }
        th { background: #333; color: #ffcc00; }
        tr:nth-child(even) { background: #1e1e1e; }
        a.button { display: block; margin: 20px auto; padding: 10px; background: #ffcc00; color: #121212; text-align: center; text-decoration: none; }
    </style>
</head>
<body>
    <h1>Orders List</h1>
    <table>
        <tr><th>Order ID</th><th>Customer ID</th><th>Quantity</th></tr>
        {{range .}}<tr><td>{{.OrderID}}</td><td>{{.CustomerID}}</td><td>{{.Quantity}}</td></tr>{{end}}
    </table>
    <a class="button" href="/">Create Order</a>
</body>
</html>

`
