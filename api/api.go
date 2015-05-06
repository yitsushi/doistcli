package api

import (
	"bytes"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const apiEndpoint string = "https://todoist.com/API/v6/sync"

var apiToken string

func SetToken(token string) {
	apiToken = token
}

type basicAPIAction struct {
	Command string            `json:"type"`
	UUID    string            `json:"uuid"`
	TempId  string            `json:"temp_id"`
	Args    map[string]string `json:"args"`
}

func (this *basicAPIAction) Send() {
	this.UUID = getRandomUUID()
	this.TempId = getRandomUUID()

	commands := []*basicAPIAction{this}

	command, _ := json.Marshal(commands)

	log.Printf("Send data: %s\n", command)
	this.httpPostRequest(string(command))
}

func (this *basicAPIAction) httpPostRequest(command string) {
	postData := url.Values{}
	postData.Set("token", apiToken)
	postData.Add("commands", command)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiEndpoint, bytes.NewBufferString(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	// resp, err := http.PostForm(apiEndpoint, postData)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Printf("%s", body)
}

func getRandomUUID() string {
	randomId, _ := uuid.NewV4()
	return randomId.String()
}
