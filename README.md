# Go Simple Backend with MySQL

This project is a simple RESTful backend API built in Go, using the Gorilla Mux router and MySQL for persistent storage. It demonstrates core backend concepts such as CRUD operations, routing, database integration, and testing in Go.

## Features

- **CRUD API for Products**:  
  - Create, Read, Update, and Delete products with fields: `id`, `name`, `quantity`, and `price`.
- **MySQL Integration**:  
  - Uses a MySQL database for persistent storage of product data.
- **RESTful Endpoints**:  
  - Follows REST conventions for endpoint design.
- **Testing**:  
  - Includes comprehensive unit tests for all endpoints using Go’s `testing` package and `httptest`.
- **Docker Support**:  
  - Includes a `docker-compose.yml` for easy setup of the MySQL database.

## Project Structure

```
go-simple-backend-mysql/
│
├── api.go           # API endpoint definitions and handler registration
├── constants.go     # Application constants (e.g., DB credentials)
├── docker-compose.yml
├── go.mod / go.sum  # Go module files
├── handlers.go      # HTTP handler functions for product operations
├── main.go          # Application entry point
├── main_test.go     # Unit tests for API endpoints
├── models.go        # Product model and DB logic
```

## Endpoints

- `GET    /products`        - List all products
- `GET    /products/{id}`   - Get a product by ID
- `POST   /products`        - Create a new product
- `PUT    /products/{id}`   - Update a product
- `DELETE /products/{id}`   - Delete a product

## How to Run

1. **Start MySQL with Docker Compose**  
   In the `go-simple-backend-mysql` directory:
   ```
   docker-compose up -d
   ```

2. **Run the Go Application**  
   ```
   go run main.go
   ```

3. **Run Tests**  
   ```
   go test
   ```

## What I Learned

- Setting up a Go project with modules and dependencies.
- Building RESTful APIs using Gorilla Mux.
- Connecting Go applications to a MySQL database.
- Implementing and running unit tests for HTTP handlers.
- Handling JSON requests and responses in Go.

## Example Request

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name":"chair","quantity":4,"price":29.99}'
```
