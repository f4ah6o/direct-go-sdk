# Testing Utilities

このディレクトリには、`direct-go`のテストに使用するユーティリティが含まれています。

## Mock Server

`mock_server.go`は、MessagePack RPC プロトコルをシミュレートするモック WebSocket サーバーを提供します。

### 使用例

```go
import "github.com/f4ah6o/direct-go-sdk/direct-go/testutil"

func TestSomething(t *testing.T) {
    // モックサーバーを作成
    mockServer := testutil.NewMockServer()
    defer mockServer.Close()
    
    // RPCメソッドのレスポンスを設定
    mockServer.OnSimple("get_me", map[string]interface{}{
        "id": "user123",
        "display_name": "Test User",
    })
    
    // クライアントを作成してモックサーバーに接続
    client := direct.NewClient(direct.Options{
        Endpoint: mockServer.URL(),
    })
    client.Connect()
    defer client.Close()
    
    // テスト実行
    user, err := client.GetMeWithContext(context.Background())
    // assertions...
}
```

### 機能

- **OnSimple**: 定数値を返すシンプルなハンドラーを登録
- **OnError**: エラーを返すハンドラーを登録
- **On**: カスタムロジックを含むハンドラーを登録
- **SendNotification**: サーバープッシュ通知をシミュレート
- **GetReceivedMessages**: 受信したRPCリクエストの履歴を取得（アサーション用）

## テストの実行

```bash
# 全テストを実行
go test ./...

# verboseモードで実行
go test -v ./...

# 特定のテストのみ実行
go test -run TestClientConnect

# カバレッジを確認
go test -cover ./...
```

## 現在のテストカバレッジ

`client_test.go` には以下のテストが含まれています：

- `TestClientConnect` - WebSocket接続とセッション作成
- `TestClientCallRPC` - 基本的なRPC呼び出し
- `TestClientRPCError` - エラーハンドリング
- `TestGetMeWithContext` - context対応のユーザー取得API
- `TestSendTextWithContext` - context対応のメッセージ送信API
