package models

type Owner struct {
	Login string `json:"login"`
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
	Count int `json:"total_count"`
	Items []Item
}