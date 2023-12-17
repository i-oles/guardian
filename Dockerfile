FROM arm64v8/golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o guardian .

EXPOSE 8080

CMD ["./guardian"]