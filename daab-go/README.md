# daabgo

Go言語版の daab (direct agent assist bot) フレームワーク

## インストール

```bash
go install github.com/f4ah6o/daabgo/cmd/daabgo@latest
```

## 使い方

### ボットの初期化

```bash
mkdir mybot && cd mybot
daabgo init
```

### ログイン

```bash
daabgo login
```

### ボットの実行

```bash
daabgo run
```

## ボットスクリプトの書き方

```go
package main

import (
    "github.com/f4ah6o/daabgo/bot"
)

func main() {
    robot := bot.New()
    
    // pingに応答
    robot.Respond("ping", func(res bot.Response) {
        res.Send("PONG")
    })
    
    // すべてのメッセージを受信
    robot.Hear(".*", func(res bot.Response) {
        fmt.Printf("Heard: %s\n", res.Text())
    })
    
    robot.Run()
}
```

## ライセンス

MIT License