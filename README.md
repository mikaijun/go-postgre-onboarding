# このリポジトリに関して
下記解説記事で作成したリポジトリになります

https://qiita.com/miumi/items/6a70d5b9a1a61cc39dc8

# 手順

## 1: cloneする
```
git clone https://github.com/mikaijun/go-postgre-onboarding.git
```

## 2: VSCodeを開き。左下の青部分をクリック
![スクリーンショット 2023-10-09 8 27 45](https://github.com/mikaijun/go-postgre-onboarding/assets/74134232/79c1b98c-4688-42e3-a063-1e24b8765728)

## 3: 「コンテナーで再度開く」をクリック
![スクリーンショット 2023-10-09 8 27 58](https://github.com/mikaijun/go-postgre-onboarding/assets/74134232/1ceb9598-d3e7-4ae0-b514-8648d90b79c0)

## 4: 動作確認

![スクリーンショット 2023-10-09 8 37 31](https://github.com/mikaijun/go-postgre-onboarding/assets/74134232/b19ec492-82e2-4113-9b99-cc8fcc46c13d)

goファイルを実行
```zsh
vscode ➜ /workspaces/go-postgres (master) $ go run main.go
```

以降別タブのターミナルで実行

getの確認
```zsh
vscode ➜ /workspaces/go-postgres/api (master) $ curl http://localhost:8080/users/get
```
```zsh
[{"Id":1,"Name":"Smith"},{"Id":2,"Name":"Johnson"},{"Id":3,"Name":"Brown"}]
```

postの確認
```zsh
vscode ➜ /workspaces/go-postgres/api (master) $ curl -X POST   -d '{ "name": "Bob" }' http://localhost:8080/users/create
```
```zsh
{"Id":4,"Name":"Bob"}
```

再度getの確認
```zsh
vscode ➜ /workspaces/go-postgres/api (master) $ curl http://localhost:8080/users/get
```
```zsh
[{"Id":1,"Name":"Smith"},{"Id":2,"Name":"Johnson"},{"Id":3,"Name":"Brown"},{"Id":4,"Name":"Bob"}]
```
