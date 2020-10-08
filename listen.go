package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleGhostWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var u ghostMemberUpdate
	err = json.Unmarshal(body, &u)
	if err != nil {
		return
	}

	if u.Member.Current.Email == "" {
		return
	}

	for _, l := range u.Member.Current.Labels {
		if l == "API" {
			return
		}
	}

	fmt.Printf("[Ghost -> SIB] %s\n", u.Member.Current.Email)
	createSibContact(sib, u.Member.Current.Email)
}

func handleSendinblueWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var u sendinblueMemberUpdate
	err = json.Unmarshal(body, &u)
	if err != nil {
		return
	}

	if len(u.ListID) == 0 || u.ListID[0] == cfg.SendinblueListID {
		return
	}

	if u.Event != "list_addition" {
		return
	}

	err = gh.CreateMember(u.Email)
	result := "V"
	if err != nil {
		result = "X"
	}

	fmt.Printf("[SIB -> Ghost] %s [%s]\n", u.Email, result)
}

type ghostMemberUpdate struct {
	Member struct {
		Current struct {
			Email  string   `json:"email"`
			Labels []string `json:"labels"`
		} `json:"current"`
	} `json:"member"`
}

type sendinblueMemberUpdate struct {
	Email  string  `json:"email"`
	Event  string  `json:"event"`
	ListID []int64 `json:"list_id"`
}
