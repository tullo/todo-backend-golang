FROM golang:1.13.7-alpine3.11 as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o todo-backend .

FROM alpine:3.11
ENV ALLOWED_ORIGINS=*
ENV PORT=8080
COPY --from=builder /go/src/app/todo-backend /usr/local/bin/
ENTRYPOINT ["todo-backend"]
