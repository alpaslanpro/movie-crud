FROM golang:1.24.1-alpine3.21

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o app .

EXPOSE 8080

CMD ["./app"]