# TODO List

A simple TODO list Rest API build wilt Gin, Gorm and SQLite.

Test database (test.db) is also provided for testing purpose.

## Installation

sqlite3 is required for this project.


```go mod download```

## Run

```
go build
./SK-todo
```


## Document
Swagger API Document generated by go-swagger.
Access http://localhost:8080/swagger/index.html

### Structure
```
SK-todo
│  go.mod
│  go.sum
│  main.go --------------> Entry point                                            
│  readme.md
│  test.db --------------> Test database
│  
├─.idea
│      .gitignore
│      SK-todo.iml
│      
├─api -------------------> API handler
│      common.go --------> Common handler / Error Response
│      task.go ----------> Task handler (Team and User)
│      team.go ----------> Team handler (Team management)
│      user.go ----------> User handler (User management)
│      
├─auth
│      jwt.go -----------> WT service
│      
├─docs ------------------> Swagger document
│      docs.go
│      swagger.json
│      swagger.yaml
│      
├─middleware            
│      middleware.go -----> Middleware, including CORS and Auth
│      
├─model ------------------> Database model
│      init.go -----------> Database init
│      migration.go ------> Database migration
│      task.go -----------> Task model
│      team.go -----------> Team model
│      user.go -----------> User model
│      
├─operation --------------> CRUD operation
│      task.go              
│      team.go
│      user.go
│      
├─resp-code --------------> Response code and message
│      code.go
│      msg.go
│      
├─router
│      router.go ---------> Router
│      
├─serializer -------------> JSOn serializer    
│      common.go
│      response.go
│      task.go
│      team.go
│      user.go
```




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
- [ ] Real-time collaboration

### Authentication (either by team or user)
- [x] JWT token
- [x] User auth by token
- [x] Team affairs permission check