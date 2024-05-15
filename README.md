# 日経 デジタル新卒研修 2024 Web の基礎

## ディレクトリ/アプリケーション構成

- `proxy/`: proxy server。http://localhost:3001 で起動する悪質な proxy。OS やブラウザの設定でこの proxy を使うように設定すると、**HTTP 通信のレスポンス結果を勝手に書き換える**。
- `simple-website`: HTML, CSS, JS を持つシンプルな Web サイト。 http://localhost:8082 で起動する。
- `tcp-server/`: シンプルな HTTP/1.1 を実装した TCP サーバ。http://localhost:3000 で起動する。

## 起動方法

Docker が必要

### すべてのアプリを起動

```bash
make dev_all
```

or

```bash
docker compose build && docker compose up
```

### 前編用のアプリを起動

```bash
make dev_first
```

or

```bash
docker compose build tcp_server proxy_server simple_website && docker compose up tcp_server proxy_server simple_website
```
