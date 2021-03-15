package main

import (
	"fmt"
	"log"
	"os"
)

func isFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func printNet(n *net, site *Site, done *map[string]void, res *string, depth int) {
	if site == nil {
		return
	}
	if _, exist := (*done)[site.Url]; exist {
		return
	}
	for i := 0; i < depth; i++ {
		*res += "> "
	}
	*res += site.Title + " " + site.Url + "\n"
	(*done)[site.Url] = v
	for link := range site.Links {
		if _, exist := n.Sites[link]; !exist {
			continue
		}
		printNet(n, n.Sites[link], done, res, depth + 1)
	}
}

func outputMarkdown(n *net, path string) {
	fmt.Println("Markdown outputting...")
	var file *os.File
	var err error

	//pathUrl, err := url.Parse(path)
	//if err != nil {
	//	log.Println("Err: path parse error, ", err)
	//	return
	//}
	//fileName, _ := url.Parse("output.md")
	//path = pathUrl.ResolveReference(fileName).String()
	path += "/output.md"

	if isFileExist(path) {
		file, err = os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0666)
	} else {
		file, err = os.Create(path)
	}
	if err != nil {
		log.Println("Err: file open or create failed, ", err)
		return
	}
	defer file.Close()
	var output = ""
	done := make(map[string]void)
	for _, site := range n.Sites {
		printNet(n, site, &done, &output, 0)
	}
	_, err = file.WriteString(output)
	if err != nil {
		log.Fatalln("Fatal: file write error, ", err)
	}
	fmt.Println(output)
	fmt.Println("Markdown file has been saved at: " + path)
}
