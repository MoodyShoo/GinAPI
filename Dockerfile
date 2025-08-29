FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod  go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o api ./cmd/main.go

FROM debian:bullseye-slim
WORKDIR /root/
COPY --from=builder /app/api .

EXPOSE 8080
ENTRYPOINT ["./api"]
