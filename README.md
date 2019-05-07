# study-golang
Go研修用リポジトリ

# 必要なもの
- direnv
- (DB接続が必要な場合)Docker For MacでDBコンテナが起動している


※lenetのMy環境が構築されていれば満たしています


# 準備
以下のリポジトリを自分のリポジトリへforkする

https://github.com/WHITEPLUS/study-golang

自身の環境に合わせてローカル環境を構築する

```bash
cd $GOPATH/src/(自分のGithubユーザー名)
git clone (forkした自身のリポジトリ)
vim .envrc
# 以下の内容に変更
# export GITHUB_USER=自分のGithubユーザー名
```

# コンテナ起動
```make start```のコマンドでコンテナを起動する

次のURLにアクセスできるか確認する
- http://localhost:32744/
    - hello assets index が表示される
- http://localhost:32744/wep/systems.health.check
    - hc が表示される

# ディレクトリ構成
```
study-golang
├── asset - 静的コンテンツのトップルート
├── chap1 - 演習1 work space
├── chap2 - 演習2 work space
├── chap3 - 演習3 work space
├── chap4 - 演習4 work space
└── kubernetes - 演習環境
```
