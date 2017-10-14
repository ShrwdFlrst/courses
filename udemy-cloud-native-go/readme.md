Cloud Native Go
---------------

## Curl commands
    curl -v localhost:9090/api/books
    curl -v -L -H "Content-Type: application/json" -X POST -d '{"title":"Infinite Jest","author":"David Foster Wallace","isbn":"7944656987"}' localhost:9090/api/books
    curl -v localhost:9090/api/books/4543456000
    curl -vL -X DELETE localhost:9090/api/books/4543456000
    curl -v -L -H "Content-Type: application/json" -X PUT -d '{"title":"Infinite Jest","author":"David Foster Wallace","isbn":"7944656987"}' localhost:9090/api/books/7944656987
