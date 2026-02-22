# Wiki

このディレクトリは、`wails-chatapp` のナレッジと仕様をまとめる場所です。

## 目次

- [Backend / Specs / MVP仕様](./backend/specs/mvp-spec.md)
- [Backend / Architecture / Clean Architecture設計](./backend/architecture/clean-architecture.md)
- [Backend / Planning / 実装段取り](./backend/planning/implementation-steps.md)
- [Backend / Operations / ブランチ情報](./backend/operations/branches.md)
- [Backend / Operations / WebSocket無停止更新戦略](./backend/operations/websocket-rollout-strategy.md)

## 運用ルール（簡易）

- 仕様変更時は先にこの`wiki/`を更新する
- 実装の判断基準は`wiki/`を正とする
- 大きな方針変更はPR単位で履歴を残す
- `wiki/backend/<genre>/*.md` / `wiki/frontend/<genre>/*.md` 形式で管理する
  - 例: `wiki/backend/architecture/clean-architecture.md`
