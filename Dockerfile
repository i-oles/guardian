FROM arm32v7/golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app/guardian

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]