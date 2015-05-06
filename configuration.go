package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

var configurationFile string

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

	conf.checkApiTokenIsExists()

	return &conf, nil
}

func (this *Configuration) checkApiTokenIsExists() {
	if this.APIToken != "" {
		return
	}

	fmt.Print("API Token: ")
	fmt.Scanf("%s", &this.APIToken)

	this.save()
}

func (this *Configuration) save() {
	jsonContent, err := json.Marshal(this)
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
