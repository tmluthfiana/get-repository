package model

import (
	"encoding/json"
)

type Owner struct {
	Login string `json:"lofin"`
}

type Item struct {
	ID              int `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Owner           Owner 
	Description     string `json:"description"`
	CreatedAt       string `json:"created_at"`
	StargazersCount int    `json:"stargazers_count"`
}

type JSONData struct {
	Count int `json:"count"`
	Items []Item
}