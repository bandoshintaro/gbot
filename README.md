# gbot

issueを作ると自動でorganizationに招待してくれるやつ

## 使い方

### webhookの設定

- githubの設定画面でissueイベントにチェックをつけた上で、このurlでwebhookを設定する

```
[hostname]:8080/github/organization-invite-webhook
```

### botに使うtokenを作成

- organizationへの権限を持つアカウントでpersonal access tokenを作る

### botを起動

使用できる環境変数

```
GITHUB_ORGANIZATION_NAME # 招待するorganizationの名前
HUBOT_GITHUB_API         # 連携するGithubAPIのURL
HUBOT_GITHUB_TOKEN       # 作成したpersonal access token
```

実行例
```
HUBOT_GITHUB_TOKEN='myapitoken' GITHUB_ORGANIZATION_NAME='my-og' HUBOT_GITHUB_API='https"//hoge.jp' ./bin/hubot
```

