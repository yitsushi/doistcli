package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

var configurationFile string

// Configuration is the data structure of the configuration file
type Configuration struct {
	APIToken string `json:"api_key"`
}

func parseConfiguration() (*Configuration, error) {
	if configurationFile == "" {
		configurationFile = getUserHome() + "/.doistcli.conf"
	}

	conf := Configuration{}

	jsonContent, err := ioutil.ReadFile(configurationFile)

	if err != nil {
		log.Println("Configuration file does not exist.")
	} else {
		err = json.Unmarshal(jsonContent, &conf)

		if err != nil {
			return nil, err
		}
	}

	conf.checkAPITokenIsExists()

	return &conf, nil
}

func (c *Configuration) checkAPITokenIsExists() {
	if c.APIToken != "" {
		return
	}

	fmt.Println("You can find your API Token under your account on todoist.com")
	fmt.Println("  Settings -> Todoist Settings -> Account -> API token")
	fmt.Print("API Token: ")
	fmt.Scanf("%s", &c.APIToken)

	c.save()
}

func (c *Configuration) save() {
	jsonContent, err := json.Marshal(c)
	if err != nil {
		panic(fmt.Sprintf("Can not encode the configuration file. (%s)\n", err))
	}

	err = ioutil.WriteFile(configurationFile, jsonContent, 0600)

	if err != nil {
		panic(fmt.Sprintf("Can not save the configuration file. (%s)\n", err))
	}
}

func getUserHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}
