package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type void struct {}
var v void

type Site struct {
	Url   string
	Title string
	Desc  string
	Links map[string]*Site
}

type SiteInfo struct {
	Title string
	Links map[string]void
}

func GetSiteInfo(urlString string) *SiteInfo {
	info := GetInfo()

	// Request HTML page
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Println("Err: http get error, ", err)
		return nil
	}
	req.Header.Add("User-Agent", info.UserAgent)
	client := http.Client {
		Timeout: info.Timeout * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Err: http request error, ", err)
		return nil
	}
	defer res.Body.Close()

	// Convert to utf-8 encoded HTML
	// https://www.jianshu.com/p/fad896422f6e
	//bytes, err := bufio.NewReader(res.Body).Peek(1024)
	//if err != nil {
	//	log.Println("Err: html utf-8 decode error, ", err)
	//	return nil
	//}
	//encoding, _, _ := charset.DetermineEncoding(bytes, "")
	//utf8Reader := transform.NewReader(res.Body, encoding.NewDecoder())
	// https://github.com/PuerkitoBio/goquery/blob/7ebd145bd7b75771461b557dc2e53525191d3dd3/doc/tips.md
	//_, encoding, _ := charset.DetermineEncoding(bytes, "")
	//utfBody, err := iconv.NewReader(res.Body, encoding, "utf-8")
	//if err != nil {
	//	log.Println("Err: html utf-8 decode error, ", err)
	//	return nil
	//}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Err: html parse error, ", err)
		return nil
	}

	siteInfo := SiteInfo {
		Links: make(map[string]void),
	}

	// Read
	// Find title
	siteInfo.Title = doc.Find("title").Contents().Text()
	siteInfo.Title = strings.TrimSpace(strings.Trim(siteInfo.Title, "\n"))

	u, err := url.Parse(urlString)
	if err != nil {
		log.Println("Err: url parse error, ", err)
		return &siteInfo
	}
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		link, exist := selection.Attr("href")
		if exist {
			l, err := url.Parse(link)
			if err != nil {
				log.Println("Err: link's url parse error, ", err)
				return
			}
			siteInfo.Links[u.ResolveReference(l).String()] = v
			//siteInfo.Links = append(siteInfo.Links, u.ResolveReference(l).String())
		}
	})

	return &siteInfo
}
