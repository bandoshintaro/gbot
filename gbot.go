package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Webhook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	//owner := config.Organization
	token := config.AccessToken
	url := config.GithubAPI

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	// enterpriseに対応
	var client *github.Client
	if url != "" {
		client, _ = github.NewEnterpriseClient(url, url, tc)
	} else {
		client = github.NewClient(tc)
	}

	// list all repositories for the authenticated user
	repos, _, _ := client.Repositories.List(ctx, "", nil)
	fmt.Println(repos[0].Name)

	/*

	   #issueを作るとorganizationに自動招待
	   robot.router.post "/github/organization-invite-webhook", (req, res) ->

	       # https://developer.github.com/v3/activity/events/types/#issuesevent
	       data = JSON.parse(req.body.payload)

	       #open以外は何もしない
	       if data.action not in ["opened"]
	           return res.end ""

	       url = require('url')
	       issue = data.issue
	       user = issue.user
	       invited_user = user.login

	       # https://developer.github.com/v3/orgs/members/#members-list
	       member_url = "#{url_api_base}/orgs/#{owner}/members"
	       github.get member_url, (members, error) ->
	           og_member = []

	           for member, i in members
	               og_member.push member.login

	           if invited_user not in og_member
	               repo = data.repository.name
	               number = issue.number

	               # issueにコメント
	               url = "#{url_api_base}/repos/#{owner}/#{repo}/issues/#{number}/comments"
	               data = { "repo": repo, "number": number, "body":"ようこそ！" }
	               github.post url, data, (body, error) ->
	                    console.log(error)
	                    res.end ""

	               # ユーザー招待のリクエストを送信
	               url = "#{url_api_base}/orgs/#{owner}/memberships/#{invited_user}"
	               github.put url, (body, error) ->
	                    console.log(error)
	                    res.end ""

	           #既にmemberになってたら何もしない
	           else
	               res.end ""
	*/
	//js, err := json.Marshal(genes)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	//w.Write(js)
}
