# selectstamp

セレクトスタンプの送信と回答集計を行うサンプルBotです。

## 機能

- `menu` または `メニュー` コマンドでセレクトスタンプを送信
- 選択結果に応じた機能を実行:
  - **uuid占い**: ランダムなUUIDを生成して運勢を占う
  - **ミラサポplus事例表示**: 中小企業向け支援事例をランダム表示

## 使い方

```bash
# daab-go-examples ディレクトリから実行
go run ./selectstamp
```

Botに対して「メニュー」とメッセージを送信すると、セレクトスタンプが表示されます。

## 環境変数

| 変数名 | 説明 | デフォルト |
|--------|------|-----------|
| `DIRECT_ACCESS_TOKEN` | directアクセストークン | 必須 |
| `DIRECT_ENDPOINT` | WebSocketエンドポイント | 必須 |
| `DEBUG_SERVER` | デバッグサーバーURL | `http://localhost:9999` |

## セレクトスタンプAPI

このサンプルでは以下のdirect APIを使用しています：

- `send_select_stamp` - セレクトスタンプの送信
- `get_action` - 回答状況の取得
