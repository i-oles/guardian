FROM arm64v8/golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o guardian ./cmd

EXPOSE 8080

CMD ["./guardian"]