# direct-go

Go言語版の direct クライアントライブラリ

[![Go Reference](https://pkg.go.dev/badge/github.com/f4ah6o/direct-go-sdk/direct-go.svg)](https://pkg.go.dev/github.com/f4ah6o/direct-go-sdk/direct-go)
[![direct-go porting coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/f4ah6o/direct-go-sdk/main/.github/badges/direct-go-porting-coverage.json)](./COVERAGE.md)

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

## メッセージ受信

`OnMessage` の callback と `Client.Messages` channel は独立した配信先です。
`OnMessage` を登録しても `Messages` の受信を消費しません。

`Messages` は `Options.MessageChannelSize`（未指定時は100）の bounded channel です。
channel が満杯になると、受信ループは接続終了まで backpressure を適用します。
`Close` または接続断では待機中の配信をキャンセルし、実装する `MessageMetrics` に drop 理由 `connection_closed` を通知します。
再接続可能な client の lifetime channel のため、`Messages` 自体は `Close` で閉じません。
channel consumer は `Done` と組み合わせて終了を検知してください。
callback はメッセージごとに別 goroutine で実行され、panic は debug log へ記録されます。

## リリース

Git tag を使用してバージョン管理します：

```bash
git tag direct-go/v0.1.0
git push --tags
```
