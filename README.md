# おとしものアプリ (Lost-Item App)
おとしものアプリは、落とし物を撮影するだけで世界の人に共有できるツールです。このアプリを使うことで、警察に届けるまでもない落とし物を見つけた場合、すぐに共有することができます。

# アプリ
[おとしものアプリ](https://otoshimono.gpio.biz/)
こちらからご利用ください

## フロント
[おとしものアプリフロント](https://github.com/gpioblink/otoshimono-front)はこちらから
## 使用言語
このアプリのバックエンドはGoで開発されました。

## データベース
PostgreSQLを採用しました
Cloud SQLを使用しています
## フレームワーク,ライブラリ
以下のフレームワーク,ライブラリを使用しています。バージョン番号はgo.modに記載されています

### gin
GoのWebフレームワークです
採用理由としてGitHubのスター数,ドキュメントの豊富さ開発メンバーが使用経験ある点から採用しました

### gorm
GoのORMです
採用理由としてGitHubのスター数,ドキュメントの豊富さを基準に採用しました

## 開発環境の再現
```
git clone git@github.com:yuorei/lost-item-backend.git
```
GitHubからクローンします
開発環境にdockerを採用しました
```
docker compose up
```
アプリのバックエンドがローカルで立ち上がります

## デプロイ
Cloud Runにデプロイしています
### CI,CD
GitHubでmainにmergeされると自動でCI,CDが入ります
Cloud Runで設定を行いました
- 継続的デプロイを編集をします
- リポジトリとブランチを選択します
- Cloud Run でシークレット,環境変数をセットしてください