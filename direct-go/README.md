# direct-go

Go言語版の direct クライアントライブラリ

[![Go Reference](https://pkg.go.dev/badge/github.com/f4ah6o/direct-go-sdk/direct-go.svg)](https://pkg.go.dev/github.com/f4ah6o/direct-go-sdk/direct-go)

## バージョン

- **direct-go**: v0.1.0
- Based on: direct-js (L is B internal)

## インストール

```bash
go get github.com/f4ah6o/direct-go-sdk/direct-go
```

## 使い方

```go
package main

import (
    "fmt"
    "log"
    
    direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

func main() {
    client := direct.NewClient(direct.Options{
        Endpoint:    "wss://api.direct4b.com/albero-app-server/api",
        AccessToken: "YOUR_ACCESS_TOKEN",
    })
    
    // イベントハンドラ登録
    client.OnMessage(func(msg direct.ReceivedMessage) {
        fmt.Printf("Received: %s\n", msg.Text)
    })
    
    // 接続
    if err := client.Connect(); err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    // メッセージ送信
    client.SendText("room-id", "Hello!")
    
    // 待機
    select {}
}
```

## リリース

Git tag を使用してバージョン管理します：

```bash
git tag direct-go/v0.1.0
git push --tags
```
