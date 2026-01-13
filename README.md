# Record Store API

A RESTful API built with Go for managing a record store application. This API provides user authentication, album management, and user management functionality.

## Features

- **User Authentication**: JWT-based authentication with secure password hashing
- **User Management**: Register and manage user accounts
- **Album Management**: Create, read, update, and delete albums
- **Protected Endpoints**: Secure endpoints requiring JWT authentication
- **PostgreSQL Database**: Robust persistent data storage with GORM ORM
- **Docker Support**: Complete Docker and Docker Compose setup for easy deployment

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework
- **Database**: PostgreSQL with [GORM](https://gorm.io/) - ORM for database operations
- **Authentication**: JWT (JSON Web Tokens)
- **Language**: Go 1.25.2
- **Containerization**: Docker & Docker Compose
- **Environment Management**: godotenv

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

- Docker and Docker Compose installed
- Or Go 1.25.2 or higher (for local development)

### Quick Start with Docker

1. Clone the repository:
```bash
git clone https://github.com/SteliosKoulinas/recordStoreApi.git
cd recordStoreApi
```

2. Copy the environment template:
```bash
cp .env.example .env
```

3. Copy and configure the environment file:
```bash
cp .env.example .env
```

Then edit `.env` with your database credentials and JWT secret. 

4. Build and run with Docker Compose:
```bash
docker-compose up --build
```

The API will be available at `http://localhost:3187`

**First time setup note:** The database containers and volumes will be created automatically on first run.

### Stopping the Application

```bash
# Stop all containers
docker-compose down

# Stop and remove volumes (clean database)
docker-compose down -v
```

### Local Development Setup

1. Clone and navigate to the project:
```bash
git clone https://github.com/SteliosKoulinas/recordStoreApi.git
cd recordStoreApi
```

2. Create `.env` file with your local database configuration

3. Install dependencies:
```bash
go mod download
```

4. Run the application:
```bash
go run main.go
```

### Setup

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

## Authentication Examples

### Register a new user
```bash
curl -X POST http://localhost:3187/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "yourpassword"
  }'
```

### Login
```bash
curl -X POST http://localhost:3187/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "yourpassword"
  }'
```

### Create an album (Protected)
```bash
curl -X POST http://localhost:3187/api/albums \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{
    "artist": "The Beatles",
    "title": "Abbey Road",
    "year": 1969
  }'
```

Protected endpoints require a valid JWT token in the `Authorization` header:
```
Authorization: Bearer <your_jwt_token>
```

## Environment Variables

### Required Variables
- `DATABASE_URL` - Full PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT token signing

### Optional Variables
- `PORT` - Server port (default: 3187)
- `GIN_MODE` - `debug` or `release` (default: debug)
- `ENVIRONMENT` - `development` or `production` (default: development)

**⚠️ Security:** Never commit actual credentials to `.env`. Use `.env.example` as template and configure with real values locally.

## Database

The application uses PostgreSQL for data persistence. Tables for `Album` and `Users` models are automatically created on startup using GORM's `AutoMigrate`.

When running with Docker Compose:
- PostgreSQL runs in the `recordstore` service
- Database volume is persisted in `postgres_data`
- Database is automatically initialized based on `DATABASE_URL`

## Docker

### Docker Compose Commands

```bash
# Build and start all services
docker-compose up --build

# Start services (without rebuild)
docker-compose up

# Start in background
docker-compose up -d

# Stop all containers
docker-compose down

# Stop and remove volumes (deletes database data)
docker-compose down -v

# View logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f api
docker-compose logs -f recordstore
```

### Build Docker Image Manually
```bash
docker build -t recordstore-api .

# Run with environment variables
docker run -p 3187:3187 \
  -e DATABASE_URL="your-database-url" \
  -e JWT_SECRET="your-secret-key" \
  recordstore-api
```

## Dependencies

Key dependencies include:
- `github.com/gin-gonic/gin` - Web framework
- `github.com/golang-jwt/jwt/v5` - JWT token handling
- `gorm.io/driver/postgres` - PostgreSQL driver
- `gorm.io/gorm` - ORM framework
- `golang.org/x/crypto` - Cryptographic functions for password hashing
- `github.com/joho/godotenv` - Environment variable loading

See `go.mod` for the complete list of dependencies.

## Security Notes

- Never commit `.env` file to version control
- Use strong, randomly generated `JWT_SECRET` in production
- Change default database credentials
- Use `GIN_MODE=release` in production
- Consider using SSL/TLS for database connections in production (change `sslmode=disable`)

## License

This project is open source and available under the MIT License.

## Author

Created by [Stelios Koulinas](https://github.com/SteliosKoulinas)
