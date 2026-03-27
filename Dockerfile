FROM golang:1.21-alpine

ARG MY_VAR
ENV MY_VAR=$MY_VAR

WORKDIR /app 

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

CMD ["./main"]
