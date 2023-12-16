FROM balenalib/raspberry-pi-debian-golang:1.21.3

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app/guardian

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]