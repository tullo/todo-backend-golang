# todo-backend-golang
A backend for TodoMVC implemented with Go using no external dependencies

## Docker

build the docker container with:
```
docker build . -t todo-backend
```

run the docker container with:
```
docker run -p 8080:80 todo-backend:latest
```