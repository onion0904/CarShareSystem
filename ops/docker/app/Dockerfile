# Builder stage
FROM golang:1.24.2 AS builder

# 作業ディレクトリを設定
WORKDIR /workspace

# go.work とあれば go.work.sum をコピー
COPY go.work* ./

# モジュールディレクトリを丸ごとコピー
COPY app   ./app
COPY pkg   ./pkg

# ワークスペース全体の依存を同期
RUN go work sync

# app モジュールをビルド
WORKDIR /workspace/app
RUN go build -o /usr/local/bin/app-server cmd/server/main.go

# 実行ステージ
FROM golang:1.24.2 AS runner
# 作業ディレクトリを設定
WORKDIR /app

# ビルド成果物をコピー
COPY --from=builder /usr/local/bin/app-server /usr/local/bin/app-server

# マイグレーションファイルをコピー（コード内の相対パスに合わせて配置）
COPY app/infrastructure/db/migrations /app/infrastructure/db/migrations

# ポート公開
EXPOSE 8080

# エントリポイント
ENTRYPOINT ["app-server"]