FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/main

RUN go build -o main .

RUN chmod 751 main

EXPOSE 8080

CMD ["./main"]
