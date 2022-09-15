# Echo Golang Task Day 2

## How to Run ?

- `go run cmd/app/*`

## Routes

Method | Route | Description
-------|-----:|:----------:
GET | http://localhost:8000/v1/users | List all user
GET | http://localhost:8000/v1/users/{id} | Show detail user
POST| http://localhost:8000/v1/users | Create new user
PUT | http://localhost:8000/v1/users/{id} | Edit user
DEL | http://localhost:8000/v1/users/{id} | Delet user