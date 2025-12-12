# direct-go-sdk

[![CI](https://github.com/f4ah6o/direct-go-sdk/actions/workflows/ci.yaml/badge.svg)](https://github.com/f4ah6o/direct-go-sdk/actions/workflows/ci.yaml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> [!IMPORTANT]
> これは**非公式**のSDKです。L is B社およびdirect公式チームとは関係ありません。

direct(direct4b.com) チャットサービス用のGo SDKです。低レベルのSDK(`direct-go`)と高レベルのBotフレームワーク(`daab-go`)を提供します。

## モジュール

- **[direct-go](./direct-go)**: direct Go SDK - WebSocket/MessagePack RPCクライアント
- **[daab-go](./daab-go)**: direct-goを使用したBotフレームワークおよびCLIツール

## 参照リポジトリ

このSDKは以下の公式リポジトリを参照して開発されています：

- [lisb/direct-js](https://github.com/lisb/direct-js) - direct JavaScript SDK
- [lisb/daab](https://github.com/lisb/daab) - Direct as a Bot フレームワーク
- [lisb/hubot-direct](https://github.com/lisb/hubot-direct) - direct用Hubotアダプター

## ドキュメント

詳細な開発ドキュメントは [AGENTS.md](./AGENTS.md) を参照してください。

## ライセンス

[MITライセンス](./LICENSE)
