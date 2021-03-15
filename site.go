package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
)

type Site struct {
	Url   string
	Title string
	Desc  string
	Links []string
}

func GetSiteInfo(urlString string) *Site {
	info := GetInfo()

	// Request HTML page
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Println("Err: http get error, ", err)
		return nil
	}
	req.Header.Add("User-Agent", info.UserAgent)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Err: http request error, ", err)
		return nil
	}
	defer res.Body.Close()

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Err: html parse error, ", err)
		return nil
	}

	site := Site{ Url: urlString }
	u, err := url.Parse(urlString)
	if err != nil {
		log.Println("Err: url parse error, ", err)
		return &site
	}

	// Read
	// Find title
	site.Title = doc.Find("title").Contents().Text()

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		link, exist := selection.Attr("href")
		if exist {
			l, err := url.Parse(link)
			if err != nil {
				log.Println("Err: link's url parse error, ", err)
				return
			}
			site.Links = append(site.Links, u.ResolveReference(l).String())
		}
	})

	return &site
}