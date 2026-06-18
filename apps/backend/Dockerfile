FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]