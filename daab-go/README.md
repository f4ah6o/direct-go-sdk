# daab-go

Go言語版の daab (direct agent assist bot) フレームワーク

[![Go Reference](https://pkg.go.dev/badge/github.com/f4ah6o/direct-go-sdk/daab-go.svg)](https://pkg.go.dev/github.com/f4ah6o/direct-go-sdk/daab-go)

## バージョン

- **daab-go**: v0.1.0
- Based on: daab (L is B internal)

## インストール

### bot パッケージ (外部プロジェクトから利用)

```bash
go get github.com/f4ah6o/direct-go-sdk/daab-go
```

### CLI ツール

```bash
go install github.com/f4ah6o/direct-go-sdk/daab-go/cmd/daabgo@latest
```

## 使い方

### 外部プロジェクトからの利用

```go
package main

import (
    "context"
    "log"

    "github.com/f4ah6o/direct-go-sdk/daab-go/bot"
)

func main() {
    robot := bot.New(
        bot.WithName("mybot"),
    )

    // ping に応答
    robot.Respond("ping", func(ctx context.Context, res bot.Response) {
        res.Send("PONG")
    })

    // すべてのメッセージを受信
    robot.Hear(".*", func(ctx context.Context, res bot.Response) {
        log.Printf("Heard: %s", res.Text())
    })

    if err := robot.Run(context.Background()); err != nil {
        log.Fatal(err)
    }
}
```

### CLI を使った開発

```bash
# ログイン
daabgo login

# 実行
daabgo run
```

## リリース

Git tag を使用してバージョン管理します：

```bash
git tag daab-go/v0.1.0
git push --tags
```
