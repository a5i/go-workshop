# go-workshop

A Go-based implementation of the "RealWorld" (Conduit) application. This project serves as a full-stack example featuring a RESTful API built with Go and a frontend built with Vue.js, demonstrating common patterns for authentication, database management, and API routing.

## Features

- **User Management**: User registration and login with JWT-based authentication.
- **Article Management**: Listing articles, fetching feeds, and creating content.
- **Storage**: SQLite integration using GORM for data persistence.
- **Frontend**: A complete Vue.js frontend bundled within the static assets.
- **Seeding**: Utility to populate the database with fake data for development.

## Tech Stack

- **Backend**: Go
- **Web Framework**: [Echo](https://github.com/labstack/echo)
- **ORM**: [GORM](https://github.com/jinzhu/gorm)
- **Database**: SQLite
- **Authentication**: JWT (JSON Web Tokens)
- **Frontend**: Vue.js

## Getting Started

### Installation

1. Ensure you have Go installed on your machine.
2. Clone the repository:
   ```bash
   git clone https://github.com/a5i/go-workshop.git
   cd go-workshop/realworld
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application

To start the backend server and serve the frontend:
```bash
go run cmd/main.go
```
The server will start on port `3333`. You can access the application at `http://localhost:3333`.

### Seeding the Database

To populate your local SQLite database with sample users and articles:
```bash
go run seed/main.go
```

### Running Tests

Run the test suite to verify the API endpoints:
```bash
go test ./cmd/...
```

## API Endpoints

### Articles
- `GET /api/articles` - List all articles.
- `GET /api/articles/feed` - Get the user's personalized feed (Requires Authentication).

### Users
- `POST /api/users` - Create a new user account.
- `POST /api/users/login` - Authenticate a user and receive a JWT.
