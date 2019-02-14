# go_echo_api

## 説明
GolangのフレームワークEchoの勉強のために作ってみました。

## 起動

#### 手順

1. [Clone](#clone)
2. [GitConfigにユーザ情報を追加](#gitconfigにユーザ情報を追加)
3. [サービスの起動](#サービスの起動)

#### Clone

このプロジェクトを任意のディレクトリ配下にCloneします。
Cloneが完了したら、プロジェクトルートに移動してください。

```bash
$ git clone https://{ 自身のGitHubアカウントのユーザ名 }@github.com/sumashin/go_echo_api.git
$ cd go_echo_api
```

#### GitConfigにユーザ情報を追加

Gitのログに表示するユーザ名とメールアドレスを追加します。

```bash
$ git config --local user.name { 自身の名前(GitHubアカウントのユーザ名である必要はないです) }
$ git config --local user.email { 自身のGitHubアカウントのメールアドレス }
```

#### サービスの起動

セットアップスクリプトの実行後の処理が完了したら、以下コマンドを実行してサービスを起動します。

```bash
$ go run main.go
```
