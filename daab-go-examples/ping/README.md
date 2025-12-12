# Ping Bot Example

daab-go フレームワークを使ったシンプルなボットの例です。

## セットアップ

1. `.env` ファイルを作成し、アクセストークンを設定:
```
HUBOT_DIRECT_TOKEN=your_access_token_here
```

2. ボットを起動:
```bash
go run ./main.go
```

## 使い方

ボット名（デフォルト: `pingbot`）を含めてメンションするとコマンドに反応します。

| コマンド | 説明 | 例 |
|---------|------|-----|
| `ping` | `PONG` を返す | `@pingbot ping` |
| `echo <text>` | テキストをそのまま返す | `@pingbot echo hello` |
| `time` | サーバー時刻を返す | `@pingbot time` |
| `shout <text>` | テキストを送信する | `@pingbot shout hello everyone` |

## 注意事項

### ボット起動時の過去メッセージへの反応

ボットが起動していない間にトークで発言した内容に、起動時に反応することがあります。

**原因:**
- `start_notification` が `false` を返した場合、`reset_notification` でリセットしてから通知を開始する
- この時、未読メッセージがまとめて通知として届く

**対処方法:**
1. メッセージの `created_at` をボット起動時刻と比較し、起動前のメッセージを無視する
2. ボット起動時に過去のメッセージを既読にする（`read_message` API）

これは direct-js (元の daab) と同様の挙動です。
