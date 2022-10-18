FROM golang:1.19.0-bullseye
WORKDIR /app
COPY . .
RUN go build main.go
CMD ["./main"]