# gbot

This is application to invite a user into your organization.

## Usage

- add webhook to your organization
  - check 'issue comments'
  - add '/webook' to PayloadURL
- create your Personal access token
- start gbot

## Configuration
```
GBOT_PORT          #default: 8080
GBOT_GITHUBAPI     #if you use Github Enterprise, you should set this variable.
GBOT_ACCESSTOKEN   #your personal access token
GBOT_ORGANIZATION  #your organization
```

