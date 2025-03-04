# CarShareSystem

## 目的
- CleanArchitectureを学ぶ
- GraphQLと~SQLboiler~sqlcを使ってみる
- 友達と車をシェアしたい

## ユーザー登録
- サインイン(auth)
    -メールアドレス
    - パスワード
    - パスワードの確認
- プロフィール設定(user)
    - 名前
    - アイコン(何個かから選ぶ動物的なのから選ぶ)
- メンバーを招待(member)
    - Lineで招待
    - QRコードで招待
    - 招待リンクをコピーする
- 日程決め(calendar)
    - カレンダー形式
    - 予約作成でメモを追加できる
    - 通学で使うかどうかを追加できる
    - 一週間前(月曜日)から次の週(月曜日から日曜日)の登録可能
    - ルール
        - 重要な用事を2個+普通の用事を2個以上>=4
        - 四日まで登録可能
        - かぶった場合
        - 重要な用事と重要な用事(相談)
        - 重要な用事と普通の用事(重要な用事を優先)
        - 普通の用事と普通の用事(相談)
- メンバーリストの表示


## 使用技術
- API形式
    - GraphQL
- CI/CDパイプライン
    - Github Actions
    - Docker
- デプロイ先
    - Cloudflare
- orm
    - ~SQLBoiler~ (Maintenance Mode)
    - sqlc
- framework
    - echo
- ログイン認証
    - JWT
- SQL
    - MySQL

## ディレクトリの説明

```
├──app # アプリケーションコード
|   ├── cmd # アプリケーションのスタート地点 (main.go等)
|   ├── config # 各種設定値
|    ├── docs # APIドキュメント
|    ├── domain # ビジネスロジックの中核 、各種ドメインオブジェクト
|    ├── infrastructure # データベースや外部APIへの実装詳細
|    ├── presentation # ユーザーへの表⽰や⼊⼒ (HTTP等)
|    ├── server # HTTPサーバーの設定やルーティング
|    └── usecase # ユースケース
├── ops # 各種環境周りに関するファイル
└── pkg # ドメインロジックをもたない汎⽤的な処理
```

## 参考

- [Goで学ぶGraphQLサーバーサイド入門](https://zenn.dev/hsaki/books/golang-graphql)
- [Go ⾔語で構築するクリーンアーキテクチャ設計](https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn?productVariantID=itEzQN5gKZX8gXMmLTEXAB)