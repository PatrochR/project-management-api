# Project Management API

A RESTful API for managing projects, tasks, and users. Built with Go, PostgreSQL, and JWT authentication.

## Features

- User registration and login with secure password hashing
- JWT-based authentication and authorization
- CRUD operations for projects and tasks
- Role-based access control for project members
- Comprehensive unit tests
- Swagger/OpenAPI documentation

## Technologies Used

- Go (Golang)
- PostgreSQL
- Gorilla Mux (HTTP router)
- JWT (github.com/golang-jwt/jwt)
- bcrypt for password hashing
- Testify for unit testing
- Docker (optional for containerization)

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL
- (Optional) Docker & Docker Compose

### Clone the Repository

```bash
git clone https://github.com/yourusername/project-management-api.git
cd project-management-api
```

### Database Setup

Create the database and run migrations (if any):

```bash
psql -U postgres -c "CREATE DATABASE projectdb;"
# Run migration scripts if provided
```

### Run the Application

```bash
go run ./cmd/main.go
```

The API will start on `localhost:8888`.

### Run with Docker (Optional)

```bash
docker-compose up --build
```

## API Documentation

Swagger UI is available at:  
`http://localhost:8888/swagger/`

## Running Tests

```bash
go test ./...
```

## Example API Usage

### Register

```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

### Login

```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

### Create Project

```http
POST /api/projects
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json

{
  "name": "New Project",
  "description": "Project description"
}
```

## Project Structure

```
internal/
  entity/         # Domain models
  repository/     # Database access
  usecase/        # Business logic
  interface/      # HTTP handlers, middleware
mocks/            # Mock implementations for testing
cmd/              # Application entry point
docs/             # Swagger/OpenAPI docs
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/foo`)
5. Create a new Pull Request

## License

MIT License

---

**Contact:**  
For questions or feedback, open an issue or contact [rezacharsetad@gmail.com](mailto:rezacharsetad@gmail.com)
