package main

import "fmt"

type net struct {
	RootUrl    string
	SitesQueue [] struct {
		oriUrl string
		desUrl string
	}
	Sites      map[string]*Site
}

func (n *net) Init() {
	n.Sites = map[string]*Site{"": {Links: make(map[string]*Site)}}
	n.SitesQueue = append(n.SitesQueue, struct {
		oriUrl string
		desUrl string
	}{oriUrl: "", desUrl: n.RootUrl})
}

func (n *net) Run(searchNum int)  {
	for i := 1; i <= searchNum; i++ {
		if len(n.SitesQueue) <= 0 {
			break
		}
		linkInfo := n.SitesQueue[0]
		nextSite, exist := n.Sites[linkInfo.desUrl]
		site := n.Sites[linkInfo.oriUrl]
		if !exist {
			fmt.Println(i, "/", searchNum, ":", linkInfo.desUrl)
			siteInfo := GetSiteInfo(linkInfo.desUrl)
			nextSite = &Site {
				Url: linkInfo.desUrl,
				Links: make(map[string]*Site),
			}
			if siteInfo != nil {
				nextSite.Title = siteInfo.Title
				for link := range siteInfo.Links {
					n.SitesQueue = append(n.SitesQueue, struct {
						oriUrl string
						desUrl string
					}{oriUrl: linkInfo.desUrl, desUrl: link})
				}
			}
		}
		site.Links[linkInfo.desUrl] = nextSite
		n.Sites[linkInfo.desUrl] = nextSite
		n.SitesQueue = n.SitesQueue[1:]
	}
}
