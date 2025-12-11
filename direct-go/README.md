# direct-go

Go言語版の direct クライアントライブラリ

## インストール

```bash
go get github.com/f4ah6o/direct-go
```

## 使い方

```go
package main

import (
    "fmt"
    "log"
    
    direct "github.com/f4ah6o/direct-go"
)

func main() {
    client := direct.NewClient(direct.Options{
        Endpoint:    "wss://api.direct4b.com/albero-app-server/api",
        AccessToken: "YOUR_ACCESS_TOKEN",
    })
    
    // イベントハンドラ登録
    client.OnMessage(func(msg direct.Message) {
        fmt.Printf("Received: %s\n", msg.Text)
    })
    
    // 接続
    if err := client.Connect(); err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    // メッセージ送信
    client.Send("room-id", direct.TextMessage{Text: "Hello!"})
    
    // 待機
    select {}
}
```
