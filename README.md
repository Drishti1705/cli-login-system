# CLI Login System

A secure Command Line Interface (CLI) based authentication system developed in **Go**. This project demonstrates authentication fundamentals, secure password management, session handling, account security mechanisms, and **Multi-Factor Authentication (MFA) using TOTP** with SQLite as the persistence layer.

---

## Project Overview

The CLI Login System provides a complete authentication workflow that allows users to:

- Register new accounts
- Login securely
- Reset passwords
- Enable and verify TOTP-based two-factor authentication
- View profile information
- Logout
- Prevent unauthorized access through account lockout

The application follows a layered architecture separating business logic, database operations, and user interaction, making it scalable, maintainable, and easy to extend.

---

# Features

## Authentication

- Secure User Registration
- User Login
- Password Reset
- Session Management
- Logout
- Profile Management

---

## Multi-Factor Authentication (MFA)

The system implements **TOTP (Time-Based One-Time Password)** as an additional authentication layer.

Features include:

- Generate unique TOTP secret keys for users
- Enable two-factor authentication during registration
- Verify OTP during login
- Time-based OTP expiration handling
- Support for authenticator applications like Google Authenticator/Authy
- Additional security layer even if passwords are compromised

---

## Security

- Password hashing using **bcrypt**
- Plain-text passwords are never stored
- Duplicate username validation
- Account lockout after multiple failed login attempts
- Password verification before reset
- Secure TOTP verification
- Session-based authentication management

---

## Database

- SQLite integration
- Automatic database initialization
- User persistence
- Storage of authentication metadata including TOTP configuration

---

## Developer Features

- Repository Pattern
- Layered Architecture
- Docker Support
- Unit Tests
- Git Version Control

---

# Architecture

```
                    +----------------+
                    |      CLI       |
                    | (User Input)   |
                    +-------+--------+
                            |
                            ▼
                    +----------------+
                    | Authentication |
                    | Business Logic |
                    +-------+--------+
                            |
              +-------------+-------------+
              |                           |
              ▼                           ▼
      +---------------+          +---------------+
      | Password Auth |          | TOTP Verify   |
      |    bcrypt     |          | OTP Validation|
      +---------------+          +---------------+
              |                           |
              +-------------+-------------+
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

# Authentication Flow

```
User Registration

        |
        ▼

Create Username & Password

        |
        ▼

Generate Unique TOTP Secret

        |
        ▼

Store User Information

        |
        ▼

Configure Authenticator App

        |
        ▼

Account Created
```

### Login Flow

```
Enter Username

        |
        ▼

Enter Password

        |
        ▼

Validate Password using bcrypt

        |
        ▼

Enter TOTP Code

        |
        ▼

Verify OTP

        |
        ▼

Create Session

        |
        ▼

Access Granted
```

---

# Project Structure

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
│   │   ├── session.go
│   │   └── totp.go
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

# Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | Backend Development |
| SQLite | Database |
| bcrypt | Password Hashing |
| TOTP | Multi-Factor Authentication |
| OTP Library | OTP Generation and Verification |
| Docker | Containerization |
| Git | Version Control |

---

# Getting Started

## Clone the Repository

```bash
git clone https://github.com/<your-username>/cli-login-system.git

cd cli-login-system
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run the Application

```bash
go run cmd/main.go
```

or

```bash
go build -o app.exe cmd/main.go

./app.exe
```

---

# Run with Docker

Build and start the application:

```bash
docker compose up --build
```

---

# Running Tests

Run all tests:

```bash
go test ./...
```

Run with coverage:

```bash
go test ./... -cover
```

---

# Sample Execution

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

Generating TOTP Secret...

Configure Authenticator Application

✅ User Registered Successfully
```

### Login

```
Username: drishti

Password: ********

Enter TOTP Code: 493821

✅ Authentication Successful

Welcome drishti
```

---

# Security Implementation

The application follows authentication best practices including:

## Password Security

- Password hashing using bcrypt
- No plain-text password storage
- Secure password verification
- Password validation before sensitive operations

## Account Protection

- Failed login attempt tracking
- Account lockout mechanism
- Duplicate username prevention

## Multi-Factor Authentication

- TOTP-based second authentication factor
- Unique secret generated for each user
- Time-based OTP validation
- OTP verification after successful password authentication
- Protection against compromised passwords

## Session Security

- Session-based login management
- Protected user operations
- Secure logout functionality

---

# TOTP Implementation Details

TOTP (Time-Based One-Time Password) provides an additional authentication factor by generating temporary OTP codes based on:

- User-specific secret key
- Current timestamp
- Fixed validity interval

The implementation includes:

- Generating a unique secret during account setup
- Associating the secret with a user account
- Validating OTP during login
- Rejecting expired or invalid OTP codes
- Integrating MFA into the existing authentication flow

---

# Future Enhancements

Potential improvements include:

- JWT Authentication
- Password Strength Validation
- Email Verification
- Role-Based Access Control (RBAC)
- Audit Logging
- Backup Authentication Codes
- REST API Version
- Web-based authentication dashboard

---

# Design Decisions

This project follows a layered architecture to improve maintainability and scalability.

### CLI Layer

Handles user interaction, menu navigation, and command execution.

### Authentication Layer

Implements business rules including:

- Registration
- Login
- Password reset
- Session management
- TOTP verification

### Repository Layer

Encapsulates database operations and provides abstraction from the storage implementation.

### Database Layer

Manages SQLite connectivity, migrations, and schema initialization.

This separation of concerns makes the application easier to test, maintain, and extend with additional authentication mechanisms.

---

# Author

**Drishti Joshi**

Go | Backend Development | Authentication Systems | MFA | SQLite | Docker