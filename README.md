# Echo Golang Task Day 2

## How to Run ?

- Clone this repo
- `go run cmd/app/*`

## Routes

### User Resource

Method | Route | Description
-------|:-----|:----------
GET | http://localhost:8000/v1/users | List all user
GET | http://localhost:8000/v1/users/{id} | Show detail user
POST| http://localhost:8000/v1/users | Create new user
PUT | http://localhost:8000/v1/users/{id} | Edit user
DEL | http://localhost:8000/v1/users/{id} | Delet user

### Book Resource

Method | Route | Description
-------|:-----|:----------
GET | http://localhost:8000/v1/books | List all book
GET | http://localhost:8000/v1/books/{id} | Show detail book
POST| http://localhost:8000/v1/books | Create new book
PUT | http://localhost:8000/v1/books/{id} | Edit book
DEL | http://localhost:8000/v1/books/{id} | Delet book

## Implementation

- ✅ CRUD
- ❎ env variable
- ❎ JWT Middleware 
- ❎ Log Middleware
- ❎ Unit Test