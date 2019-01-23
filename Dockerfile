FROM golang:1.10

ENV PORT=80

COPY ./src/todo-backend ./

RUN go build -o main .
CMD [ "./main" ]