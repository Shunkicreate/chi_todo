# ベースイメージとして公式のGoイメージを使用
FROM golang:1.22.3-alpine

# 作業ディレクトリを設定
WORKDIR /app

# モジュールの初期化
RUN go mod init github.com/shunkicreate/todo-app

# 必要な依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod tidy

# Airのインストール
RUN go install github.com/cosmtrek/air@latest

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o todo-app ./cmd

# ポートを指定
EXPOSE 8080

# Airを使用してアプリケーションを実行
CMD ["air"]
