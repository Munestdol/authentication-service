# Authentication service
## A simple service for authorization authentication
## Installation
#### git clone https://github.com/Munestdol/authentication-service.git
#### cd authentication-service
#### go mod tidy
## Using
#### go run cmd/authentication-service/main.go
## API Routes
### POST /auth/login
**Request:** 
```json
{
  "username": "username",
  "password": "password"
}
```
### **Response:**
```json
{
  "token": "JWT_TOKEN"
}
```
### GET /auth/protected
#### Headers - Authorization: Bearer JWT_TOKEN
**Response:**
```json
{
    "message": "Valid token, welcome user"
}
```
## Technologies
- GoLang
- Gin Web Framework
- JWT (JSON Web Tokens)
- PostgreSQL
