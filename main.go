package main

import "fmt"

func main() {
	//info := GetInfo()
	site := GetSiteInfo("https://zh.wikipedia.org/wiki/Python")
	if site != nil {
		fmt.Println(site.Links)
	}

}

