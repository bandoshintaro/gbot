# gbot

issueを作ると自動でorganizationに招待してくれるやつ

## 使い方

### webhookの設定

- githubの設定画面でissueイベントにチェックをつけた上で、このurlでwebhookを設定する

```
[hostname]:8080/webhook
```

### botに使うtokenを作成

- organizationへの権限を持つアカウントでpersonal access tokenを作る

### botを起動

使用できる環境変数

```
GBOT_PORT          #gbotを起動するポート番号、デフォルトは8080
GBOT_GITHUBAPI     #Github Enterprise用の設定、github.comの場合は不要
GBOT_ACCESSTOKEN   #認証に使用するPersonal Access Token
GBOT_ORGANIZATION  #招待したいOrganization
```

