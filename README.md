# CLI Login System

A secure Command Line Interface (CLI) based authentication system developed in **Go**. This project demonstrates authentication fundamentals, secure password management, session handling, and clean software architecture using SQLite as the persistence layer.

---

## Project Overview

The CLI Login System provides a complete authentication workflow that allows users to:

- Register new accounts
- Login securely
- Reset passwords
- View profile information
- Logout
- Prevent unauthorized access through account lockout

The application follows a layered architecture separating business logic, database operations, and user interaction, making it scalable and maintainable.

---

## Features

### Authentication
- Secure User Registration
- User Login
- Password Reset
- Session Management
- Logout

### Security
- Password hashing using **bcrypt**
- Duplicate username validation
- Account lockout after multiple failed login attempts
- Password verification before reset

### Database
- SQLite integration
- Automatic database initialization
- User persistence

### Developer Features
- Repository Pattern
- Layered Architecture
- Docker Support
- Unit Tests
- Git Version Control

---

## Architecture

```
                    +----------------+
                    |      CLI       |
                    | (User Input)   |
                    +-------+--------+
                            |
                            ▼
                    +----------------+
                    | Authentication |
                    |  Business Logic|
                    +-------+--------+
                            |
                            ▼
                    +----------------+
                    |   Repository   |
                    | Database Layer |
                    +-------+--------+
                            |
                            ▼
                    +----------------+
                    |    SQLite DB   |
                    +----------------+
```

---

## Project Structure

```
cli-login-system
│
├── cmd
│   └── main.go
│
├── internal
│   ├── auth
│   │   ├── login.go
│   │   ├── register.go
│   │   ├── password.go
│   │   ├── reset_password.go
│   │   └── session.go
│   │
│   ├── cli
│   │   └── menu.go
│   │
│   ├── database
│   │   ├── database.go
│   │   └── migration.go
│   │
│   ├── repository
│   │   └── user_repository.go
│   │
│   └── models
│       └── user.go
│
├── tests
│
├── data
│   └── login.db
│
├── Dockerfile
├── docker-compose.yml
├── README.md
├── go.mod
└── go.sum
```

---

## Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | Backend Development |
| SQLite | Database |
| bcrypt | Password Hashing |
| Docker | Containerization |
| Git | Version Control |

---

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/<your-username>/cli-login-system.git

cd cli-login-system
```

---

### Install Dependencies

```bash
go mod tidy
```

---

### Run the Application

```bash
go run cmd/main.go
```

or

```bash
go build -o app.exe cmd/main.go

./app.exe
```

---

## Run with Docker

Build and start the application:

```bash
docker compose up --build
```

---

## Running Tests

Run all tests:

```bash
go test ./...
```

Run with coverage:

```bash
go test ./... -cover
```

---

## Sample Execution

```
========== CLI Login System ==========

1. Register
2. Login
3. Reset Password
4. Profile
5. Logout
6. Show Users
7. Exit

Choose option: 1

Username: drishti
Password: ********

✅ User Registered Successfully
```

---

## Security Implementation

The application follows authentication best practices including:

- Password hashing using bcrypt
- Plain-text passwords are never stored
- Account lockout after repeated failed login attempts
- Duplicate username prevention
- Session-based login management
- Password reset with old password verification

---

## Testing

The project includes automated tests covering:

- Password hashing
- Password verification
- User registration
- Duplicate registration
- Login validation
- Password reset
- Session management
- Account lockout

---

## Future Enhancements

Potential improvements include:

- JWT Authentication
- Multi-Factor Authentication (MFA)
- Password Strength Validation
- Email Verification
- Role-Based Access Control (RBAC)
- Audit Logging
- REST API Version

---

## Design Decisions

This project follows a layered architecture to improve maintainability and scalability.

- **CLI Layer** – Handles user interaction.
- **Authentication Layer** – Implements business rules such as login, registration, and password reset.
- **Repository Layer** – Encapsulates all database operations.
- **Database Layer** – Manages SQLite connectivity and schema initialization.

This separation of concerns makes the application easier to test, maintain, and extend.

---

## Author

**Drishti Joshi**

Go | Backend Development | Authentication Systems | SQLite | Docker