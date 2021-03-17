package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"time"
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
		*res += "  "
	}
	*res += "- " + site.Title + " " + site.Url + "\n"
	(*done)[site.Url] = v
	for link := range site.Links {
		printNet(n, n.Sites[link], done, res, depth + 1)
	}
}

func outputMarkdown(n *net, dir string) {
	fmt.Println("Markdown outputting...")
	var file *os.File
	var err error

	pathUrl, err := url.Parse(dir)
	if err != nil {
		log.Println("Err: path parse error, ", err)
		return
	}
	dir = path.Join(pathUrl.Path, time.Now().Format("2006-01-02-15-04-05") + "_output.md")
	//fileName, _ := url.Parse("./" + time.Now().Format("2006-01-02-15-04-05") + "_output.md")
	//path = pathUrl.ResolveReference(fileName).String()
	//path += "/" + time.Now().Format("2006-01-02-15-04-05") + "_output.md"

	if isFileExist(dir) {
		file, err = os.OpenFile(dir, os.O_CREATE | os.O_WRONLY, 0666)
	} else {
		file, err = os.Create(dir)
	}
	if err != nil {
		log.Println("Err: file open or create failed, ", err)
		return
	}
	defer file.Close()
	var output = ""
	done := make(map[string]void)
	printNet(n, n.Sites[n.RootUrl], &done, &output, 0)
	_, err = file.WriteString(output)
	if err != nil {
		log.Fatalln("Fatal: file write error, ", err)
	}
	fmt.Println(output)
	fmt.Println("Markdown file has been saved at: " + dir)
}
