# ---- build stage ----
FROM golang:1.25-alpine AS build
WORKDIR /app

# 依存だけ先に取得（キャッシュ効率化）
COPY go.mod go.sum ./
RUN go mod download

# ソースをコピーしてビルド（静的バイナリ）
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /server ./cmd/server

# ---- run stage ----
FROM alpine:3.20
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=build /server /app/server
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/app/server"]
