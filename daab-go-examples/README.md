# daab-go-examples

`daab-go` を使用したサンプルBotアプリケーション集です。

## サンプル一覧

| サンプル | 説明 |
|---------|------|
| [ping](./ping) | シンプルなping/pong応答Bot |
| [selectstamp](./selectstamp) | セレクトスタンプの送信と回答集計 |
| [n8n-proxy](./n8n-proxy) | directメッセージをn8n webhookに転送するプロキシBot |
| [teams-bridge](./teams-bridge) | directとMicrosoft Teamsを連携するブリッジBot |

## セットアップ

### 1. 認証情報の設定

```bash
cd daab-go-examples
cp .env.example .env
```

`.env` ファイルを編集して認証情報を設定するか、`daabgo login` コマンドを使用してください。

### 2. サンプルの実行

```bash
go run ./ping
```

## 必要な環境変数

| 変数名 | 説明 |
|--------|------|
| `DIRECT_ACCESS_TOKEN` | directアクセストークン |
| `DIRECT_ENDPOINT` | WebSocketエンドポイントURL |

各サンプル固有の環境変数については、各サンプルのディレクトリ内のREADMEを参照してください。
