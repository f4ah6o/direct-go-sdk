# Direct4B Porting Coverage Tool

自動的にdirect-jsからdirect-goへの移植状況を追跡・レポートするツールです。

## 機能

- **RPC メソッドレベルでの比較**: direct-jsとdirect-goのRPCメソッド呼び出しを比較
- **カテゴリ別の分析**: 13のカテゴリに分類して詳細な分析
- **複数の出力形式**: Markdown、JSON、テキスト形式に対応
- **CI統合準備済み**: GitHub Actionsなどで簡単に利用可能

## インストール

```bash
# direct-goディレクトリで
cd tools/coverage
go build -o coverage
```

または、直接実行：

```bash
go run .
```

## 使い方

### 基本的な使用

```bash
# Markdownレポートを標準出力に表示
go run ./tools/coverage

# または、tools/coverage ディレクトリから
cd tools/coverage
go run .
```

### 各種オプション

```bash
# Markdownファイルに出力
go run ./tools/coverage -output COVERAGE.md

# JSON形式で出力
go run ./tools/coverage -format json -output coverage.json

# テキストサマリーを表示
go run ./tools/coverage -format text

# 詳細ログを表示
go run ./tools/coverage -verbose

# パスを明示的に指定
go run ./tools/coverage -js-path ../direct-js -go-path ../..

# ハードコードされたベースラインを使用（JSファイル読み込みなし）
go run ./tools/coverage -use-baseline
```

### コマンドラインオプション

| オプション | デフォルト | 説明 |
|-----------|-----------|------|
| `-js-path` | `../direct-js` | direct-jsディレクトリへのパス |
| `-go-path` | `../..` | direct-goディレクトリへのパス |
| `-output` | (stdout) | 出力ファイルパス。指定しない場合は標準出力 |
| `-format` | `markdown` | 出力形式: `json`, `markdown`, `text` |
| `-verbose` | `false` | 詳細なログを表示 |
| `-use-baseline` | `false` | JSソースを読まずにハードコードされたベースラインを使用 |
| `-version` | - | バージョン情報を表示 |

## 出力形式

### Markdown

人間が読みやすい形式で、以下の情報を含みます：

- サマリー（カバレッジ率、実装済み/未実装メソッド数）
- カテゴリ別のカバレッジ表
- 各カテゴリの詳細（実装済み・未実装メソッドのリスト）
- 優先度付き推奨事項

### JSON

機械可読形式で、以下の構造：

```json
{
  "metadata": {
    "generated_at": "2025-12-10T...",
    "tool_version": "1.0.0",
    "js_path": "/path/to/direct-js",
    "go_path": "/path/to/direct-go"
  },
  "summary": {
    "total_js_methods": 82,
    "total_go_methods": 8,
    "coverage_percentage": 9.76,
    "implemented_count": 8,
    "missing_count": 74
  },
  "categories": [ ... ],
  "all_methods": { ... }
}
```

### Text

コンソール表示用の簡潔なサマリー：

```
Direct4B Porting Coverage: 9.76% (8/82 methods)

Top 3 Categories by Coverage:
  🟡 Session & Auth: 57.1%
  🟠 Domain Management: 28.6%
  🟠 Talk/Room Management: 22.2%
```

## カテゴリ分類

ツールは82のRPCメソッドを以下の13カテゴリに分類します：

1. **Session & Auth** (7) - セッション管理と認証
2. **User Management** (11) - ユーザー情報の管理
3. **Domain Management** (7) - ドメイン/組織の管理
4. **Department Management** (3) - 部署階層の管理
5. **Talk/Room Management** (9) - トーク/ルームの管理
6. **Message Operations** (17) - メッセージ送受信・検索
7. **File & Attachment Management** (6) - ファイルアップロード・ダウンロード
8. **Note Management** (6) - ノート機能
9. **Announcement Management** (4) - お知らせ機能
10. **Push Notification Management** (2) - プッシュ通知設定
11. **Conference/Call Management** (5) - ビデオ/音声通話
12. **Miscellaneous** (5) - その他の機能

## カバレッジステータス

各カテゴリには視覚的なステータスが表示されます：

- 🟢 **80%以上** - 良好なカバレッジ
- 🟡 **50-79%** - 中程度のカバレッジ
- 🟠 **20-49%** - 低いカバレッジ
- 🔴 **20%未満** - 非常に低いカバレッジ

## CI統合

### GitHub Actions

```yaml
name: Coverage Report

on:
  pull_request:
  push:
    branches: [main]

jobs:
  coverage:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Generate Coverage Report
        run: |
          go run ./tools/coverage \
            -format markdown \
            -output COVERAGE.md

      - name: Upload Report
        uses: actions/upload-artifact@v3
        with:
          name: coverage-report
          path: COVERAGE.md
```

### PRコメント自動投稿

```yaml
      - name: Comment on PR
        if: github.event_name == 'pull_request'
        uses: actions/github-script@v6
        with:
          script: |
            const fs = require('fs');
            const report = fs.readFileSync('COVERAGE.md', 'utf8');
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: report
            });
```

## 仕組み

### 1. RPCメソッド抽出

**JavaScript (direct-js):**
- `direct-js/lib/direct-node.js` から正規表現で `.call("method_name")` パターンを抽出
- 82メソッドをハードコードされたベースラインと照合

**Go (direct-go):**
- `direct-go/**/*.go` から `c.call("method_name")` および `c.Call("method_name")` を抽出
- 現在8メソッドを実装

### 2. カバレッジ計算

```
Coverage = (Go実装済みメソッド数 / JS全メソッド数) × 100%
```

### 3. カテゴリ別分析

各メソッドを機能カテゴリに分類し、カテゴリごとのカバレッジを算出します。

## メンテナンス

### ベースラインの更新

direct-jsのAPIバージョンが上がり、新しいメソッドが追加された場合：

1. ツールを verbose モードで実行して新しいメソッドを検出：
   ```bash
   go run ./tools/coverage -verbose
   ```

2. `baseline.go` を編集して新しいメソッドを適切なカテゴリに追加

3. `categoryOrder` と `jsMethodsByCategory` を更新

### カテゴリの追加・変更

1. `baseline.go` の `jsMethodsByCategory` マップを編集
2. `categoryOrder` スライスに新しいカテゴリを追加
3. 必要に応じて `categorizeMethod()` 関数を更新

## トラブルシューティング

### "no such file or directory" エラー

パスが正しいか確認してください：

```bash
# 現在地を確認
pwd

# 相対パスを調整
go run ./tools/coverage -js-path ../direct-js -go-path .
```

### カバレッジが0%と表示される

Go のソースファイルが見つからない可能性があります：

```bash
# verbose モードで確認
go run ./tools/coverage -verbose
```

### 期待と異なるメソッド数

ベースラインを使用してみてください：

```bash
go run ./tools/coverage -use-baseline
```

## 開発者向け

### ファイル構成

```
tools/coverage/
├── main.go       # CLIエントリポイント
├── baseline.go   # 82 JSメソッドの定義
├── extractor.go  # ソースコードからのメソッド抽出
├── analyzer.go   # カバレッジ計算とカテゴリ分析
├── reporter.go   # JSON/Markdown/テキスト出力
└── README.md     # このファイル
```

### テスト

```bash
# ユニットテストを追加する場合
go test ./...

# ベンチマーク
go test -bench=. ./...
```

## ライセンス

このツールはdirect-goプロジェクトの一部です。

## バージョン履歴

- **v1.0.0** (2025-12-10) - 初回リリース
  - 82 JSメソッドの追跡
  - 13カテゴリでの分類
  - JSON/Markdown/テキスト出力対応
