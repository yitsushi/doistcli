package api

import (
	"encoding/json"
	"log"
	//	"strings"
)

const apiEndpoint string = "https://todoist.com/API/v6/sync"

type basicAPIAction struct {
	Command string            `json:"type"`
	UUID    string            `json:"uuid"`
	TempId  string            `json:"temp_id"`
	Args    map[string]string `json:"args"`
}

func (this *basicAPIAction) Send() {
	command, _ := json.Marshal(this)

	log.Printf("Send data: %s\n", command)
}
