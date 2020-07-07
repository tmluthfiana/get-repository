package main

import (
	"get-repository/modules"
)

func main() {
	url := "https://api.github.com/search/repositories?q=user:tmluthfiana&sort=stars&order=desc"
	modules.GetListRepositories(url)
}
