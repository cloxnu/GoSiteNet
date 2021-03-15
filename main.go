package main

import "fmt"

func main() {
	info := GetInfo()

	n := net{ Sites: make(map[string]*Site) }
	n.SitesQueue = append(n.SitesQueue, info.Url)
	run(&n)

	for _, site := range n.Sites {
		fmt.Println(site.Url, site.Title)
	}
	outputMarkdown(&n, info.OutputDir)

}

