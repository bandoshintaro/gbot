package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
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
	if hc.Event == "issues" {
		event := github.IssueCommentEvent{}
		if err := json.Unmarshal(hc.Payload, &event); err != nil {
			log.Error(err)
		}

		issue := event.GetIssue()
		username := issue.GetUser().GetLogin()
		repo := event.GetRepo().GetName()
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
		ismember, _, _ := client.Organizations.IsMember(ctx, owner, username)
		if ismember {
			log.Info(username, " is already member of ", owner)
		} else {

			tmp := "Organizationに招待しました"
			state := "admin"
			comment := &github.IssueComment{Body: &tmp}
			_, _, err = client.Issues.CreateComment(ctx, owner, repo, number, comment)
			if err != nil {
				log.Error(err)
			}
			_, _, err = client.Organizations.EditOrgMembership(ctx, username, owner, &github.Membership{State: &state})
			if err != nil {
				log.Error(err)
			}
			issue_state := "closed"
			_, _, err = client.Issues.Edit(ctx, owner, repo, number, &github.IssueRequest{State: &issue_state})
			if err != nil {
				log.Error(err)
			}

		}
	}
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
