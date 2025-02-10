# Lema Backend v1.0
## Project Overview
This project is a minature blog platform where existing users can create posts and view posts created by others. This repository contains the backend code for the platform. The backend is built using Go and Gin. It connects to a SQLite database is managed using GORM.

## Technologies Used
- Go: Backend runtime environment
- Gin: Web framework for Go
- GORM: SQL ORM for Go
- SQLite: Database for storing user information

# Directory Structure
```angular2html
lema-be/
├── .gitignore                      # Gitignore file
├── .github                         # Github Actions files for CI/CD
    └── workflows/                  
        └── deploy.yml              # Workflow File for deploying to EC2 instance
├── common                          # Utilities used across mutiple modules
├── db                              # database setup logic and useful db interfaces
├── handlers                        # members for handling http requests
├── models                          # Struct definitions for app models
└── validation/                     # Logic and utilities for data validation
```

## Installation
### Prerequisites
Ensure you have the following installed on your machine:
- Go (v1.23.x)

### Steps
1. Clone the repository
```bash 
git clone https://github.com/patrickishaf/lema-be.git && cd lema-be
```

2. Install dependencies
```bash
go mod tidy
```

3. Start the application
```bash
go run main.go
```

4. The application should now be running on your localhost and on port 8080

---

## API Endpoints
### Users Routes

- GET /users
- GET /users/{userId}
- GET /users/count

### Post Routes
- GET /posts?userId={userId}
- POST /posts
- DELETE /posts/{id}

## Database Schema
The database schema consists of two tables: `game_results`, and `game_characters`.

## Running Tests
Tests are configured to pass if there are no tests present. To run the tests, use the following command:
```bash
go test -v ./...
```
