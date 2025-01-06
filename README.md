# POS Microservice

A Point of Sale (POS) microservices system built with Go, designed to manage orders efficiently with a dark-themed POS UI.

## Features

- **Order Management:** Create, list, and manage customer orders seamlessly.
- **Dark Themed UI:** User-friendly interface tailored for POS systems.
- **Microservices Architecture:** Decoupled services for scalability and maintainability.
- **gRPC Communication:** Efficient inter-service communication using gRPC.

## Project Structure

```
go-pos-ms/
├── protobuf/                  # Protocol Buffers definitions
│   └── orders.proto           # Order service definitions
├── services/
│   ├── common/                # Shared utilities and generated protobuf code
│   │   ├── genproto/          # Generated protobuf code
│   │   └── util.go            # Utility functions
│   ├── kitchen/               # Kitchen service
│   │   ├── http.go            # HTTP handlers
│   │   └── main.go            # Service entry point
│   └── orders/                # Orders service
│       ├── handler/
│       │   └── orders/
│       │       ├── grpc.go    # gRPC handlers
│       │       └── http.go    # HTTP handlers
│       ├── service/           # Business logic
│       ├── types/             # Service-specific types
│       │   └── types.go       # Type definitions
│       ├── grpc.go            # gRPC server setup
│       ├── http.go            # HTTP server setup
│       └── main.go            # Service entry point
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies
└── README.md                  # Project documentation
```

## Technical Overview

### Microservices Architecture

The system is composed of distinct services:

- **Orders Service:** Manages order-related operations, offering both gRPC and HTTP endpoints.
- **Kitchen Service:** Handles kitchen-specific tasks, interfacing with the Orders service as needed.

Each service operates independently, promoting modularity and ease of maintenance.

### gRPC Integration

Inter-service communication is facilitated through gRPC, ensuring high performance and type safety. The `orders.proto` file defines the service contracts, which are compiled into Go code residing in the `services/common/genproto` directory.

### Dark Themed UI

The Orders service provides a web interface to display and create orders. The UI is designed with a dark theme, enhancing usability in POS environments. The HTML template is rendered by the Orders service's HTTP server.

## Getting Started

### Prerequisites

- Go 1.20 or higher

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/SIDDHARTH-PADIGAR/Go-POS-Microservices.git
   cd Go-POS-Microservices
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

### Running the Services

1. **Start the Orders Service:**

   ```bash
   cd services/orders
   go run main.go
   ```

2. **Start the Kitchen Service (optional):**

   ```bash
   cd services/kitchen
   go run main.go
   ```

### Accessing the UI

Once the Orders service is running, access the POS dashboard at `http://localhost:8080`.

## API Endpoints

- **Orders Service:**
  - `POST /orders`: Create a new order.
  - `GET /`: Fetch all orders, displayed in the POS UI.

## Technologies Used

- Go
- gRPC
- Protocol Buffers
- HTML & CSS (Dark Themed UI)

## Contributing

Contributions are welcome! Feel free to fork the repository, make enhancements, and submit a pull request.

## License

This project is open source; no specific license has been applied.

---

For more details, visit the [GitHub repository](https://github.com/SIDDHARTH-PADIGAR/Go-POS-Microservices).

```

Feel free to further customize this `README.md` to suit your project's needs! 
