# Record Store API

A RESTful API built with Go for managing a record store application. This API provides user authentication, album management, and user management functionality.

## Features

- **User Authentication**: JWT-based authentication with secure password hashing
- **User Management**: Register and manage user accounts
- **Album Management**: Create, read, update, and delete albums
- **Protected Endpoints**: Secure endpoints requiring JWT authentication
- **SQLite Database**: Lightweight persistent data storage with GORM ORM

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework
- **Database**: SQLite with [GORM](https://gorm.io/) - ORM for database operations
- **Authentication**: JWT (JSON Web Tokens)
- **Language**: Go 1.25.2

## Project Structure

```
recordStoreApi/
├── db/              # Database connection and initialization
├── handlers/        # HTTP request handlers
│   ├── album.go    # Album endpoints
│   ├── auth.go     # Authentication endpoints
│   └── users.go    # User management endpoints
├── middleware/      # Custom middleware
│   └── auth.go     # JWT authentication middleware
├── models/          # Data models
│   ├── album.go    # Album model
│   └── users.go    # User model
├── services/        # Business logic
│   ├── album_service.go  # Album operations
│   └── user_service.go   # User operations
├── utils/           # Utility functions
│   ├── jwt.go      # JWT token generation and validation
│   └── password.go # Password hashing utilities
├── main.go
└── go.mod
```

## Installation

### Prerequisites

- Go 1.25.2 or higher

### Setup

1. Clone the repository:
```bash
git clone https://github.com/SteliosKoulinas/recordStoreApi.git
cd recordStoreApi
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The API will start on `http://127.0.0.1:3187`

## API Endpoints

### Authentication
- `POST /api/register` - Register a new user
- `POST /api/login` - Login and receive JWT token

### Albums (Public)
- `GET /api/albums` - Get all albums

### Users (Public)
- `GET /api/users` - Get all users

### Albums (Protected - Requires JWT)
- `POST /api/albums` - Create a new album
- `GET /api/album/:id` - Get a specific album
- `PUT /api/album/:id` - Update an album
- `DELETE /api/album/:id` - Delete an album

## Authentication

Protected endpoints require a valid JWT token to be included in the `Authorization` header:

```
Authorization: Bearer <your_jwt_token>
```

Obtain a token by logging in or registering via the `/api/login` or `/api/register` endpoints.

## Environment

The API runs on port `3187` by default. You can modify this in `main.go`:

```go
r.Run("127.0.0.1:3187")
```

## Database

The application uses SQLite for data persistence. The database file is automatically created and initialized on startup with the `Album` and `Users` models.

## Dependencies

Key dependencies include:
- `github.com/gin-gonic/gin` - Web framework
- `github.com/golang-jwt/jwt/v5` - JWT token handling
- `gorm.io/driver/sqlite` - SQLite driver
- `gorm.io/gorm` - ORM framework
- `golang.org/x/crypto` - Cryptographic functions for password hashing

See `go.mod` for the complete list of dependencies.

## License

This project is open source and available under the MIT License.

## Author

Created by [Stelios Koulinas](https://github.com/SteliosKoulinas)
