# pistatium_blog_go

https://pistatium.dev


このブログのソースコードです。

* サーバー: App Engine + Go 1.12
* データベース: Cloud Datastore (Firestore)
* フロント: Vue.js + Vue Router (SPA)
* CD: Circle CI

App Engine を使う理由
* [Qiita: App Engine を布教したくて Go + Datastore の開発環境を Docker Compose でシュッと立ち上げた話](https://qiita.com/kimihiro_n/items/5d373440acc48488a837)


## 開発環境セットアップ

```shell
docker-compose up -d

export DATASTORE_EMULATOR_HOST=0.0.0.0:8059
export PROJECT_ID=local-app
export DATASTORE_PROJECT_ID=local-app
python3 init_local.py

cd front; 
yarn
yarn run serve
```
