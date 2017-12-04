package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/go-github/github"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Webhook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	hc, err := ParseHook(r)
	if err != nil {
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if hc.Event == "issue" {
		event := github.IssueCommentEvent{}
		if err := json.Unmarshal(hc.Payload, &event); err != nil {
			log.Error(err)
		}

		issue := event.GetIssue()
		//username := issue.GetUser().GetName()
		repo := issue.GetRepository().GetName()
		number := issue.GetNumber()

		owner := config.Organization
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
		tmp := "テスト"
		comment := &github.IssueComment{Body: &tmp}
		_, _, err = client.Issues.CreateComment(ctx, owner, repo, number, comment)
		if err != nil {
			log.Error(err)
		}
	}
	/*
		ismember, _, _ := client.Organizations.isMember(ctx, owner, username, nil)

		if ismember {
			log.Info(username, "is already member of", owner)
			return nil
		}
		_, _, err = client.Organizations.EditOrgMembership(ctx, username, owner, &github.Membership)
		if err != nil {
			log.Error(err)
			return nil
		}
	*/

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

	//w.Write(js)
}

type HookContext struct {
	Event   string
	Id      string
	Payload []byte
}

func ParseHook(req *http.Request) (*HookContext, error) {
	hc := HookContext{}

	if hc.Event = req.Header.Get("x-github-event"); len(hc.Event) == 0 {
		return nil, errors.New("No event!")
	}

	if hc.Id = req.Header.Get("x-github-delivery"); len(hc.Id) == 0 {
		return nil, errors.New("No event Id!")
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	hc.Payload = body
	if err != nil {
		return nil, err
	}

	return &hc, nil
}
