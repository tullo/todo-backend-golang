# todo-backend-golang

A backend for TodoMVC implemented with Go using no external dependencies

## Build Docker Image

Build the docker image with:

`docker build . -t todo-backend-golang`

## Run Docker Container

Run the docker container with:

`docker run -p 8080:80 todo-backend-golang`

Or use the pre-built image form DockerHub

`docker run -p 8080:80 tullo/todo-backend-golang`

## CRUD Tests

### Read (GET) response

`curl -v http://localhost:8080/todos`

```console
< HTTP/1.1 200 OK
< Access-Control-Allow-Methods: GET, POST, PATCH, DELETE
< Content-Type: application/json; charset=UTF-8

[]
```

### Create (POST)

`curl -X POST -H 'Content-Type: application/json' -d '{"title":"Foo Bar"}' http://localhost:8080/todos/`

### Update (PATCH)

`curl -X PATCH -H 'Content-Type: application/json' -d '{"completed":true}' http://localhost:8080/todos/1`

### Delete (DELETE)

`curl -X DELETE http://localhost:8080/todos/1`
