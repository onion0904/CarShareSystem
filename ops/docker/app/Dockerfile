# # Builder stage
# FROM golang:1.24.2-alpine AS builder

# # gitをインストール
# RUN apk add --no-cache git

# # 作業ディレクトリを設定
# WORKDIR /workspace

# # モジュールディレクトリを丸ごとコピー
# COPY . .

# # 依存関係を取得
# RUN go mod download

# # appディレクトリの内容を確認
# RUN ls -la app/cmd/server/

# # goのバージョンと環境を確認
# RUN go version && go env

# # app モジュールをビルド
# RUN go build -buildvcs=false -o /usr/local/bin/app-server ./app/cmd/server/main.go

# # 実行ステージ
# FROM golang:1.24.2-alpine AS runner

# # 作業ディレクトリを設定
# WORKDIR /app

# # ビルド成果物をコピー
# COPY --from=builder /usr/local/bin/app-server /usr/local/bin/app-server

# # マイグレーションファイルをコピー（コード内の相対パスに合わせて配置）
# COPY app/infrastructure/db/migrations /app/infrastructure/db/migrations

# # ポート公開
# EXPOSE 8080

# # エントリポイント
# ENTRYPOINT ["app-server"]

# Builder stage - 依存関係をシンプルに
FROM golang:1.24.2-alpine AS builder

# 必要なパッケージをインストール
RUN apk add --no-cache git

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルだけをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# main.goの内容を表示
RUN cat app/cmd/server/main.go

# ビルド（全てのフラグを使用）
RUN CGO_ENABLED=0 GO111MODULE=on go build -v -buildvcs=false -o /usr/local/bin/app-server ./app/cmd/server/main.go

# 実行ステージ
FROM alpine:latest

# 必要なランタイム依存関係
RUN apk --no-cache add ca-certificates

WORKDIR /app

# ビルド成果物をコピー
COPY --from=builder /usr/local/bin/app-server /usr/local/bin/app-server

# マイグレーションファイルをコピー
COPY app/infrastructure/db/migrations /app/infrastructure/db/migrations

# ポート公開
EXPOSE 8080

# ヘルスチェックを追加（koyebでの問題解決のため）
HEALTHCHECK --interval=5s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

# エントリポイント
ENTRYPOINT ["/usr/local/bin/app-server"]