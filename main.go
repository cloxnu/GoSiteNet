package main

import "fmt"

func main() {
	info := GetInfo()

	n := net{ RootUrl: info.Url }
	n.Init()
	n.Run(info.SearchNum)

	for _, site := range n.Sites {
		fmt.Println(site.Url, site.Title)
	}
	outputMarkdown(&n, info.OutputDir)

}

