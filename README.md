# go-pos-ms

A Point of Sale (POS) microservice system built using Go. This project provides an order management system with a simple POS-themed dark UI for listing and creating orders.

## Features
- **Order Management:** Create and list customer orders.
- **Dark Themed UI:** Simple POS-friendly design.
- **Microservices:** Separated concerns for better scalability.
- **gRPC Integration:** Order management implemented with gRPC communication.

## Project Structure
```
/go-pos-ms
├── protobuf/                  # Protobuf definitions
├── services/
│   ├── common/               # Utility functions and protobuf generation
│   ├── orders/               # Orders service with HTTP and gRPC handlers
│   └── kitchen/              # Kitchen service
├── go.mod                    # Go module definition
├── Makefile                  # Build and run commands
└── README.md
```

## Getting Started
### Prerequisites
- Go 1.20+

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-pos-ms.git
   cd go-pos-ms
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Project
Start the server:
```bash
cd services/orders
go run main.go
```

Access the POS dashboard at `http://localhost:8080`.

## API Endpoints
- `POST /orders`: Create a new order.
- `GET /`: Fetch all orders in a dark-themed POS UI.

## Technologies Used
- Go
- gRPC
- Protocol Buffers
- HTML & CSS (Dark Themed UI)

## Contributing
Contributions are welcome! Fork the repository, make changes, and submit a pull request.

## License
This project is open source, with no specific license applied yet.

