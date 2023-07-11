# monorepo

個人開発用に作成したモノレポのテンプレートです。
バックエンドは golang、モバイルアプリを react-native で作成しています。

# 使い方

github-actions 経由で ECR へ自動 push を行うため、以下の環境変数をgithub上に設定してください。

```
AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
AWS_ECR_REPO_NAME=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```
