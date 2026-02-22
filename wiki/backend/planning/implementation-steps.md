# 実装段取り（最小構成からステップアップ）

## 目的

- 最小構成で早く動くものを作り、段階的に品質と機能を上げる
- `wiki/backend/specs/mvp-spec.md` と `wiki/backend/architecture/clean-architecture.md` に沿って実装する

## 全体方針

- 各フェーズで「動作確認可能な状態」を作る
- 1フェーズごとに `wiki` を更新して仕様と実装差分をなくす
- 先にインメモリで成立させ、後で永続化へ置き換える

## Phase 0: 土台準備（最小）

- 対象:
  - `.gitignore` 整備
  - `wiki` の仕様確定（b1, b2, b4）
  - `internal/` ディレクトリ骨組み作成
- 完了条件:
  - `wails dev` で既存アプリが起動する
  - 実装ルールが `wiki` に明文化されている

## Phase 1: バックエンド最小チャット（インメモリ）

- 対象:
  - `domain/chat`, `domain/user` の最小エンティティ定義
  - `usecase/chat` に `join`, `leave`, `send` の最小ユースケース実装
  - `infrastructure/repository/memory` にリポジトリ実装
  - `interface/websocket` に Hub/Handler の最小実装
- 完了条件:
  - 複数クライアント接続でメッセージが全員に配信される
  - 切断時にコネクションがリークしない

## Phase 2: フロントエンド最小UI

- 対象:
  - ユーザー名入力
  - メッセージ一覧
  - 入力欄 + 送信ボタン
  - WebSocket接続管理（接続/切断/再接続の最小）
- 完了条件:
  - 2クライアント以上で相互にリアルタイム送受信できる
  - 参加/退出通知を表示できる

## Phase 3: Clean Architecture整合性の強化

- 対象:
  - 依存方向の見直し（`outer -> inner`）
  - DTO/Portの分離整理
  - `interface` と `infrastructure` の責務境界を明確化
- 完了条件:
  - `domain`/`usecase` が WebSocket実装詳細に依存していない
  - Repository実装が `infrastructure/repository` に統一されている

## Phase 4: テスト追加（壊れにくくする）

- 対象:
  - `usecase` のユニットテスト
  - 主要ハンドラの結合テスト（可能な範囲）
  - バリデーション系ケース（空文字、長文、異常イベント）
- 完了条件:
  - 主要ユースケースの正常系/異常系テストが揃っている
  - 最低限の回帰確認が自動でできる

## Phase 5: 永続化（SQLite）へステップアップ

- 対象:
  - `infrastructure/repository/sqlite` 実装
  - メッセージ履歴の保存/取得
  - 起動時の依存切替（memory or sqlite）
- 完了条件:
  - アプリ再起動後に履歴が参照できる
  - Repository差し替えだけでユースケースを変更しない

## Phase 6: 拡張機能

- 候補:
  - 複数ルーム
  - 認証
  - 既読/タイピング
  - ファイル送信
- 完了条件:
  - 仕様更新（`wiki`）と実装が一致している
  - 主要動作のテストが追従している

## 進め方ルール

- 1フェーズ = 1PRを基本にする
- PRには以下を必ず含める
  - 仕様差分（必要なら`wiki`更新）
  - 実装差分
  - 動作確認手順
  - 未対応事項（次フェーズ送り）
