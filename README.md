# monorepo

個人開発用に作成したモノレポのテンプレートです。
バックエンドは golang、モバイルアプリを react-native で作成しています。

# 使い方

github-actions 経由で ECR/expo へ自動 push を行うため、以下の環境変数を github 上に設定してください。

```
AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
AWS_ECR_REPO_NAME=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
EXPO_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

# dev/stg/prod 環境

## dev 環境

ローカルの開発環境のこと。

## stg 環境

dev 環境で作成したアプリケーションを手元でテストするための環境

- バックエンド：stg 環境用に構築した lambda で実行される（予定）。
- フロントエンド: expo アプリを通じて検証する。githubActions に設定した`expo update`によって mobile 層に変更があれば自動で更新される。

⚠ フロントエンドの環境変数は githubActions 上で設定する必要がある。

## prod 環境

実際のユーザーが使用している環境

- バックエンド: prod 環境用に構築した lambda で実行される。
- フロントエンド:testFlight にて実機テストを行い、問題がなければリリースする。

⚠ フロントエンドの環境変数は expo 上で設定する必要がある。
