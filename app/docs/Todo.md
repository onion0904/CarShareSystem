## 実装の順番

- domain
- usecase
- infra
- config
- presen
- server
- cmd
- docs

## todo
- authの確認
    - resolver内にあるsignup,inを踏まえてjwtを使った認証をmiddleware/auth.goに実装
- passwordをハッシュ化して保存するようにするbcryptとかを使う
- docker-composeでdbの操作権限を設定する
- E2E test
- CI/CD