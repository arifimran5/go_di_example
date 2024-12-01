# Go Dependency Example

## Features

- **Manual Dependency Management**: This project uses a simple manual dependency injection pattern, allowing for easy testing and flexibility in swapping implementations.
- **CRUD Operations**: Provides endpoints to create, read, list, and manage products.
- **SQLite Database**: Utilizes SQLite for data storage with a simple schema for products.
- **Logging**: Implements logging using a custom logger interface, supporting multiple logging strategies.

## Project Structure
```markdown
project-root/
├── cmd/
│   └── main.go              # Entry point for the HTTP API
├── handlers/                 # HTTP handlers for product operations
│   └── product_handler.go
├── services/                 # Business logic for product management
│   └── product_service.go
└── stores/                   # Data access layer for products
│   └── product_store.go
├── pkg/                          # Public packages (e.g., logger)
│   └── logger/
│       ├── default_logger.go
│       ├── logrus_logger.go
│       └── logger.go             # Logger interface
├── db/                           # Database connection and initialization
│   └── db.go                     # SQLite connection and table creation
├── go.mod                        # Go module file for dependency management
├── go.sum                        # Go module checksum file
└── README.md                     # Project documentation

```

## Getting Started

### Prerequisites

- Go 1.16 or later
- SQLite (included via `github.com/mattn/go-sqlite3`)
- Ensure CGO is enabled (for go-sqlite3 lib) by setting the environment variable:
   ```bash
     export CGO_ENABLED=1
   ```
### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/arifimran5/go_di_example.git
   cd go_di_example
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the application (Linux):

   ```bash
   go build -o bin/main ./cmd/main.go
   ```

### Running the Application (Linux)

To run the application, execute the following command:

```bash
./bin/main
```

The API will start on `http://localhost:8080`.

### API Endpoints

- **List Products**
   - `GET /products`
   - Returns a list of all products.

- **Get Product by ID**
   - `GET /products/{id}`
   - Returns the product with the specified ID.

- **Create Product**
   - `POST /products`
   - Creates a new product. The request body should be in JSON format:
     ```json
     {
       "name": "Product Name",
       "price": 19.99
     }
     ```

### Example Usage

You can use `curl` or any API client (like Postman) to interact with the API. Here are some examples:

- **List Products**:
  ```bash
  curl -X GET http://localhost:8080/products
  ```

- **Get Product by ID**:
  ```bash
  curl -X GET http://localhost:8080/products/1
  ```

- **Create Product**:
  ```bash
  curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name": "New Product", "price": 29.99}'
  ```
