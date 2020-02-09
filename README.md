# todo-backend-golang

A backend for TodoMVC implemented with Go using no external dependencies

## Build Docker Image

Build the docker image with:

`DOCKER_BUILDKIT=1 docker build -t todo-backend-golang .`

## Run Docker Container

Run the docker container with:

`docker run --hostname todomvc.go -p 8080:8080 todo-backend-golang`

Or use the pre-built image form DockerHub

`docker run --hostname todomvc.go -p 8080:8080 tullo/todo-backend-golang`

```console
$ cat /etc/hosts
...
127.0.1.1       todomvc.go
192.168.0.126   todomvc.vue
```

## CRUD Tests

### Read (GET) response

`curl -v http://todomvc.go:8080/todos`

```console
< HTTP/1.1 200 OK
< Access-Control-Allow-Methods: OPTIONS, GET, POST, PATCH, DELETE
< Content-Type: application/json; charset=UTF-8
...
[]
```

### Create (POST)

```console
curl -X POST -H 'Content-Type: application/json' \
    -d '{"title":"Foo"}' http://todomvc.go:8080/todos/
```

### Update (PATCH)

```console
curl -X PATCH -H 'Content-Type: application/json' \
    -d '{"title":"Foo","completed":true}' http://todomvc.go:8080/todos/1
```

### Delete (DELETE)

`curl -X DELETE http://todomvc.go:8080/todos/1`
