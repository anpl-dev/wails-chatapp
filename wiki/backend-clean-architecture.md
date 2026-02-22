# バックエンド設計（Go / Clean Architecture）

## 方針

- 依存方向を `outer -> inner` に固定する
- `domain` はフレームワーク非依存にする
- `usecase` は業務ロジックと入出力境界を担当する
- WebSocketやWailsなどの詳細実装は外側レイヤーに置く
- Repository実装は `infrastructure/repository` に集約する

## 想定ディレクトリ構成

```text
.
├── cmd/
│   └── wails-chatapp/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── chat/
│   │   │   ├── entity.go
│   │   │   ├── value_object.go
│   │   │   └── repository.go
│   │   └── user/
│   │       ├── entity.go
│   │       └── repository.go
│   ├── usecase/
│   │   └── chat/
│   │       ├── send_message.go
│   │       ├── join_room.go
│   │       ├── leave_room.go
│   │       ├── dto.go
│   │       └── port.go
│   ├── interface/
│   │   ├── websocket/
│   │   │   ├── handler.go
│   │   │   ├── hub.go
│   │   │   └── presenter.go
│   │   └── wails/
│   │       └── bindings.go
│   └── infrastructure/
│       ├── repository/
│       │   ├── memory/
│       │   │   ├── chat_repository.go
│       │   │   └── user_repository.go
│       │   └── sqlite/
│       │       ├── chat_repository.go
│       │       └── user_repository.go
│       └── logger/
│           └── logger.go
├── frontend/
├── wiki/
└── go.mod
```

## レイヤー責務

- `internal/domain`:
  - エンティティ、値オブジェクト、リポジトリインターフェース
  - ビジネスルールの中核
- `internal/usecase`:
  - ユースケース単位の処理
  - 入出力DTO、外部依存ポート
- `internal/interface`:
  - WebSocketの受信/配信、Wails連携
  - 外部I/Oをユースケースに橋渡し
- `internal/infrastructure`:
  - リポジトリ実装、ロガーなど詳細実装

## Repository配置ルール（確定）

- interface定義:
  - `internal/domain/*/repository.go`
- 実装:
  - `internal/infrastructure/repository/memory/*_repository.go`
  - `internal/infrastructure/repository/sqlite/*_repository.go`
- 命名:
  - `chat_repository.go`, `user_repository.go` のように集約名を揃える
- 禁止:
  - `usecase` や `interface` レイヤーにRepository実装を置かない

## WebSocket実装の分離方針

- 接続管理（Hub）は`internal/interface/websocket/hub.go`
- メッセージ受信時は`usecase`を呼び出す
- ブロードキャスト内容は`presenter`で整形
- `domain`や`usecase`は`gorilla/websocket`に依存しない

## 実装ステップ案

1. `domain`のエンティティとポートを定義
2. `usecase`に参加・退出・送信を実装
3. `interface/websocket`にHubとHandlerを実装
4. `infrastructure/repository/memory`で最小リポジトリを実装
5. Wails起動時に依存を組み立てる（DI）

## 補足

- 現在の`main.go`や`app.go`は段階的に上記構成へ移行する
- 最初はインメモリ実装で進め、永続化は後から差し替える
