package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

type user struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          bool        `json:"hireable"`
	Bio               string      `json:"bio"`
	TwitterUsername   string      `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type api struct {
	Username string
	Base     string
	User     user
}

func (ptr api) getUserInfo() user {

	gitClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}
	req, err := http.NewRequest(http.MethodGet, ptr.Base, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("user", "github")

	res, getErr := gitClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var _user user
	jsonErr := json.Unmarshal(body, &_user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return _user
}

func (ptr api) display() {
	v := reflect.ValueOf(ptr.User)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

func main() {
	var _api api
	_api.Username = "milan090"
	_api.Base = "https://api.github.com/users/" + _api.Username
	_api.User = _api.getUserInfo()
	_api.display()
}
