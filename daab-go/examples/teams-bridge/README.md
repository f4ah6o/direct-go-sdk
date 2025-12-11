# Teams Bridge Bot

direct の1:1トーク（ペアトーク）でのメッセージを n8n 経由で Microsoft Teams チャンネルに転送し、Teams からの返信を direct に返すブリッジボットです。

## ユースケース

ITサポート窓口として、direct ユーザーからの問い合わせを Teams のサポートチームに集約できます。

## アーキテクチャ

```
direct ユーザー ⇄ Bot ⇄ n8n ⇄ Microsoft Teams
```

- 1 ユーザー = 1 Teams スレッド（n8n で永続化）
- 初回問い合わせ時に Teams スレッド作成、以降は返信

## セットアップ

### 1. 環境変数

`.env` ファイルを作成:

```
HUBOT_DIRECT_TOKEN=your_direct_access_token
N8N_WEBHOOK_URL=http://n8n.local:5678/webhook/direct-to-teams
CALLBACK_PORT=8080
```

### 2. n8n ワークフロー

#### ワークフロー 1: direct → Teams

1. **Webhook** トリガー
   - HTTP Method: POST
   - Path: `/direct-to-teams`

2. **永続化ノード** (例: Redis, Google Sheets など)
   - `userId` で `threadId` を検索

3. **条件分岐**
   - スレッドあり → Microsoft Teams Reply
   - スレッドなし → Microsoft Teams Create Message → 永続化

#### ワークフロー 2: Teams → direct

1. **Microsoft Teams Trigger** (New Channel Message)

2. **永続化ノード**
   - `threadId` で `userId`, `talkId` を検索

3. **HTTP Request**
   - URL: `http://bot-host:8080/webhook/teams-reply`
   - Method: POST
   - Body:
     ```json
     {
       "userId": "...",
       "talkId": "...",
       "message": "Teams からの返信内容"
     }
     ```

### 3. Bot 起動

```bash
go run ./main.go
```

## API リファレンス

### Bot → n8n (Webhook)

```json
POST /webhook/direct-to-teams
{
  "userId": "direct_user_id",
  "talkId": "direct_talk_id", 
  "message": "問い合わせ内容"
}
```

### n8n → Bot (Callback)

```json
POST /webhook/teams-reply
{
  "userId": "direct_user_id",
  "talkId": "direct_talk_id",
  "message": "Teams からの回答"
}
```
