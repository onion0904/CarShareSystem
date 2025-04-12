FROM golang:1.22 AS builder

WORKDIR /app

# すべてのファイルをコピー
COPY . .

# app ディレクトリに移動してビルド
WORKDIR /app/app
RUN go mod download && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /app/myapp-437007-bdde37cabb9b.json /app/myapp-437007-bdde37cabb9b.json

EXPOSE 8080

CMD ["/app/server"]