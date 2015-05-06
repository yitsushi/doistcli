package api

import (
	//	"encoding/json"
	//"log"
	"strings"
)

// Project resource
type Project struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	Name              string `json:"name"`
	Color             int    `json:"color"`
	Indent            int    `json:"indent"`
	ItemOrder         int    `json:"item_order"`
	Collapsed         int    `json:"collapsed"`
	IsArchived        int    `json:"is_archived"`
	Shared            bool   `json:"shared"`
	ArchivedTimestamp int    `json:"archived_timestamp"`
	IsDeleted         int    `json:"is_deleted"`
	ArchivedDate      string `json:"archived_date"`
}

func (this *Project) IsGroup() bool {
	return strings.HasPrefix(this.Name, "* ")
}

// Project endpoint
type ProjectAdd struct {
	command string `json:"type"`
	tempId  string `json:"temp_id"`
	uuid    string `json:"uuid"`
	Project *Project
}

func (this *ProjectAdd) Prepare() *basicAPIAction {
	this.command = "project_add"
	this.tempId = "UUID HERE"
	this.uuid = "UUID HERE"

	return &basicAPIAction{
		Command: this.command,
		TempId:  this.tempId,
		UUID:    this.uuid,
		Args: map[string]string{
			"name": this.Project.Name,
		},
	}
}
