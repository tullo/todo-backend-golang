FROM golang:1.12-alpine3.9 as builder
WORKDIR /go/src/app
COPY ./src/todo-backend .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o todo-backend .

FROM alpine:3.9
ENV PORT=8080
COPY --from=builder /go/src/app/todo-backend /usr/local/bin/
ENTRYPOINT ["todo-backend"]
