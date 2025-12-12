# N8N Proxy Bot

n8n webhookにdirectのイベントを転送するプロキシBot。

## 環境変数

`.env`ファイルまたは環境変数で設定:

- `N8N_WEBHOOK_URL` - n8n webhookのURL（必須）
- `HUBOT_DIRECT_TOKEN` - direct認証トークン（`daabgo login`で設定済みの場合は不要）
- `DEBUG_SERVER` - デバッグサーバーURL（オプション、デフォルト: http://localhost:9999）

`.env`ファイル例:
```bash
N8N_WEBHOOK_URL=https://your-n8n-instance.com/webhook/xxxxx
HUBOT_DIRECT_TOKEN=your_token_here
```

## 実行方法

### 方法1: .envファイルを使用（推奨）

```bash
cd examples/n8n-proxy
# .envファイルを作成
echo "N8N_WEBHOOK_URL=https://your-n8n-instance.com/webhook/xxxxx" > .env
# daabgo loginで認証済みの場合、トークンは自動的にロード
go run main.go
```

### 方法2: 環境変数で直接指定

```bash
export N8N_WEBHOOK_URL="https://your-n8n-instance.com/webhook/xxxxx"
go run examples/n8n-proxy/main.go
```

## 動作

1. directの全メッセージを受信
2. n8n webhookにJSON payloadで転送
3. n8nからのレスポンスに応じてアクション実行:
   - `none` - 何もしない
   - `reply` - メッセージに返信
   - `send` - 指定したルームに送信
   - その他のaction stamp関連（TODO）

## Payload例

```json
{
  "version": "1.0",
  "eventType": "message_created",
  "timestamp": "2025-12-12T08:50:00+09:00",
  "bot": {
    "name": "n8nproxy"
  },
  "message": {
    "id": "123456",
    "talkId": "789",
    "userId": "456",
    "type": 1,
    "typeName": "text",
    "text": "hello",
    "created": 1702345678
  }
}
```

## Response例

```json
{
  "action": "reply",
  "text": "こんにちは！n8nからの返信です"
}
```
