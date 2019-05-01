package utils

//#region Header
import (
	mod "../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	GITHUB_ENDPOINT string = "https://api.github.com"
	USER_AGENT      string = "Dmitriy-Vas/go-discord-bot"
	MEDIA_TYPE      string = "application/vnd.github.v3+json"
	SUCCESS         int    = 200
	NOT_FOUND       int    = 404
	SYS_ERROR       int    = 0
)

var (
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			MaxConnsPerHost:     5,
		},
		Timeout: time.Second * 15,
	}
)
//#endregion

func send(method string, url *mod.URL) ([]byte, int) {
	token := os.Getenv("GITHUB")
	if len(token) == 0 {
		return nil, SYS_ERROR
	}

	request, _ := http.NewRequest(method, url.Link, nil)

	q := request.URL.Query()
	for key, value := range url.Queries {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()

	request.Header.Add("Authorization", "token "+token)
	request.Header.Add("Accept", MEDIA_TYPE)
	request.Header.Add("User-Agent", USER_AGENT)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, SYS_ERROR
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, SYS_ERROR
	}
	return data, response.StatusCode
}

func get(url *mod.URL) ([]byte, int) {
	return send(http.MethodGet, url)
}

func put(url *mod.URL) ([]byte, int) {
	return send(http.MethodPut, url)
}

func del(url *mod.URL) ([]byte, int) {
	return send(http.MethodDelete, url)
}

func User(name string) (*mod.User, int) {
	data, status := get(&mod.URL{Link: GITHUB_ENDPOINT + "/users/" + name, Queries: nil})
	switch status {
	case NOT_FOUND:
	case SYS_ERROR:
		return nil, status
	}

	var user *mod.User

	err := json.Unmarshal(data, &user)
	if err != nil {
		return nil, SYS_ERROR
	}

	return user, SUCCESS
}

func Projects(name string, page int) ([]*mod.Project, int) {
	queries := make(map[string]string)
	queries["per_page"] = "100"
	queries["page"] = strconv.Itoa(page)

	data, status := get(&mod.URL{Link: GITHUB_ENDPOINT + "/users/" + name + "/repos", Queries: queries})
	switch status {
	case NOT_FOUND:
	case SYS_ERROR:
		return nil, status
	}

	var projects []*mod.Project
	err := json.Unmarshal(data, &projects)
	if err != nil {
		return nil, SYS_ERROR
	}
	return projects, SUCCESS
}

func Stars(name string) (int, int) {
	user, status := User(name)
	if status != SUCCESS || user == nil {
		return 0, SYS_ERROR
	}

	var projects []*mod.Project

	pages := int(math.Floor(float64(user.Repositories/100)) + 1)
	for page := 1; page <= pages; page++ {
		projectsCached, status := Projects(name, page)
		if status != SUCCESS {
			return 0, SYS_ERROR
		}
		projects = append(projects, projectsCached...)
	}

	out := 0
	for _, project := range projects {
		out += project.Stars
	}
	return out, SUCCESS
}
