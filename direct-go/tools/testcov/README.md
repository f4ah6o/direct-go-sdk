# direct-go test coverage (runtime)

このディレクトリは、`direct-go` のコードカバレッジ（`go test` の runtime カバレッジ）をローカルで確認するための簡易ツール一式です。porting カバレッジ（tools/coverage）とは切り離して運用します。

## 前提
- このリポジトリ直下に `direct-go/` と `direct-js/` があるローカル環境を想定しています（ポーティングの開発フローに合わせた構成）。
- Go がインストール済みで、`direct-go` をビルドできること。

## 使い方（PowerShell）
`direct-go` ルートで実行:
```powershell
pwsh tools/testcov/run.ps1
```

生成物:
- `coverage.out`: `go test` のカバレッジプロファイル
- `coverage.html`: HTML レポート（ブラウザで確認）

オプション例:
```powershell
# プロファイル/HTMLの出力先を変える
pwsh tools/testcov/run.ps1 -CoverProfile tmp/cover.out -Html tmp/cover.html

# 特定パッケージだけ
pwsh tools/testcov/run.ps1 -Packages @('direct-go/client', 'direct-go/auth')
```

## CI との併用
- porting カバレッジ（tools/coverage）とは別ジョブとして `go test ./... -coverprofile=coverage.out` を回し、`go tool cover -func=coverage.out` をログ出力、`coverage.html` をアーティファクト化する構成を推奨します。
- 閾値チェックを入れる場合は `go tool cover -func=coverage.out` の `total:` 行をスクリプトで抽出して判定してください。
