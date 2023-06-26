# kotonohako-all-in
kotonohako 開発リポジトリを全て1つにまとめたものです。

# 開発に必要なツール
- docker
- asdf
- direnv
- golang
- nodejs
- yarn

開発ツールのバージョンは全て asdf で .tool-versions 以下に記してあります。

# ローカル実行
```shell
docker-compose up
```
これにより、 以下サーバーが起動します。
- localhost:3306 に DB 
- localhost:8080 に API server
- localhost:3000 に web front server

## db のみ実行したい場合
docker-compose コマンドを実行した場合、全てプロダクション環境と同じ構成でビルドが実行されるため、実行まで時間がかかり、開発に支障でる場合があります。
その場合は、 db のみ docker-compose で起動し、他サーバーは直接実行してください。

```shell
docker-compose up db
```

### API server の起動
```shell
asdf install
direnv allow
go run server.go
```

### web front server の起動
```shell
yarn
yarn dev
```


