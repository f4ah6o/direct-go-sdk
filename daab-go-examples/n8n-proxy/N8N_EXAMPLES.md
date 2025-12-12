# N8N Workflow Response Examples

n8nから返すJSONレスポンスの例です。

## 1. 何もしない (none)

```json
{
  "action": "none"
}
```

## 2. メッセージに返信 (reply)

```json
{
  "action": "reply",
  "text": "こんにちは！n8nからの返信です"
}
```

## 3. 特定のルームに送信 (send)

```json
{
  "action": "send",
  "roomId": "1613852158778671104",
  "text": "管理者に通知しました"
}
```

## 4. エラー処理の例

n8nのFunction Nodeで以下のように実装:

```javascript
// 受信データを取得
const payload = $input.all()[0].json;

// ユーザー情報をチェック
if (payload.message && payload.message.user) {
  const user = payload.message.user;
  console.log(`Message from: ${user.email || user.displayName || user.id}`);
}

// 特定のキーワードに反応
if (payload.message.text.includes("help")) {
  return {
    action: "reply",
    text: `こんにちは ${payload.message.user.displayName}さん！何かお困りですか？`
  };
}

// デフォルトは何もしない
return {
  action: "none"
};
```

## 受信するペイロード例（拡張版）

```json
{
  "version": "1.0",
  "eventType": "message_created",
  "timestamp": "2025-12-12T09:41:00+09:00",
  "bot": {
    "name": "n8nproxy"
  },
  "message": {
    "id": "123456789",
    "talkId": "1613852158778671104",
    "userId": "1613560334893711360",
    "user": {
      "id": "1613560334893711360",
      "email": "user@example.com",
      "displayName": "山田太郎",
      "name": "yamada.taro"
    },
    "type": 1,
    "typeName": "text",
    "text": "help",
    "created": 1702345678
  }
}
```
