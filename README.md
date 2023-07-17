# monorepo

個人開発用に作成したモノレポのテンプレートです。
バックエンドは golang、モバイルアプリを react-native で作成しています。

# 使い方

github-actions 経由で ECR/testFlight へ自動 push を行うため、以下の環境変数を github/expo 上に設定してください。

github☟

```
AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
AWS_ECR_REPO_NAME=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
BASE64_ENCODED_ASC_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
EXPO_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

expo☟

```
IOS_ASC_API_KEY_ID=xxxxxxxxxxxxxxxxxxxx
IOS_ASC_API_KEY_ISSUER_ID=xxxxxxxxxxxxxxxxxxxx
IOS_ASC_APP_ID=xxxxxxxxxxxxxxxxxxxx
```
