# User Service

A robust user management microservice built with Go, providing secure user authentication and management capabilities.

## Features

- User registration and authentication
- Secure password hashing using bcrypt
- JWT-based authentication
- CRUD operations for user management
- Protected routes with middleware
- PostgreSQL database integration
- Environment-based configuration

## Tech Stack

- **Language**: Go 1.24.0
- **Web Framework**: Gin
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment Variables**: godotenv

## API Routes

### Public Routes

- `POST /users` - Create a new user
  - Required fields: FirstName, LastName, EmailAddress, Password
- `POST /login` - Authenticate user and get JWT token
  - Required fields: EmailAddress, Password

### Protected Routes

- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user
- `GET /protected` - Protected route example

## Setup

1. Clone the repository
2. Create a `.env` file with the following variables:
   ```
   DB_HOST=your_db_host
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=your_db_port
   JWT_SECRET=your_jwt_secret
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run .
   ```

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Protected routes with middleware
- Environment variable configuration
- Input validation and sanitization

## Development

The project uses a Makefile for common development tasks. Available commands:

```bash
make build    # Build the application
make run      # Run the application
```
