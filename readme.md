# TODO List

A simple TODO list Rest API build wilt Gin, Gorm and SQLite.

In memory database is enabled by default for developing.

Test database (test.db) is also provided for testing purpose.

## Installation

sqlite3 is required for this project.


```go mod download```

## Run

```go run main.go```

## Document
Swagger API Document generated by go-swagger.
Access http://localhost:8080/swagger/index.html

## Features

### User Management
- [x] User registration
- [x] User login
- [x] User info update
- [x] User delete
- [x] User team list

### Team Management
- [x] Team creation
- [x] Team info update
- [x] Team delete
- [x] Team member add
- [x] Team member remove
- [x] Team member list

### Task Management (either by team or user)
- [x] Task creation
- [x] Task info update
- [x] Task delete
- [x] Task list with filter of limit, offset and status
- [x] Task order by created_at, status or other field in db. ASEC or DESC
- [x] Task Search by name or description
- [x] Task detail info
- [ ] Real-time collection

### Authentication (either by team or user)
- [x] JWT token
- [x] User auth by token
- [x] Team affairs permission check