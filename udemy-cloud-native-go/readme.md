Cloud Native Go
---------------

## Curl commands
    curl -v localhost:9090/api/books
    curl -v -L -H "Content-Type: application/json" -X POST -d '{"title":"Infinite Jest","author":"David Foster Wallace","isbn":"7944656987"}' localhost:9090/api/books
    curl -v localhost:9090/api/books/4543456000
    curl -vL -X DELETE localhost:9090/api/books/4543456000
    curl -v -L -H "Content-Type: application/json" -X PUT -d '{"title":"Infinite Jest","author":"David Foster Wallace","isbn":"7944656987"}' localhost:9090/api/books/7944656987

    # Access on docker container
    curl -v 192.168.99.100:8080/api/books

## Docker commands

    docker pull golang:1.8.4-alpine3.6
    docker build -t udemy-cloud-native-go:1.0.0 .
    docker tag udemy-cloud-native-go:1.0.0 shrwdflrst/udemy-cloud-native-go:1.0.0
    docker images
    docker login
    docker push
    docker run -it -p 8080:8080 shrwdflrst/udemy-cloud-native-go:1.0.0
