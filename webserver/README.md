# Example 9

## Running the exmaple

To run this exmaple, from the root of this project:

```sh
go run ./v9/*.go
```

Routes:

GET  localhost:8080/  
GET  localhost:8080/todos  
POST localhost:8080/todos  
GET  localhost:8080/todos/{todoId}

To add a todo:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
