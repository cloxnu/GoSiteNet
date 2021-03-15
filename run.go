package main

import "fmt"

func run(n *net) {
	info := GetInfo()

	for i := 1; i <= info.SearchNum; i++ {
		if len(n.SitesQueue) <= 0 {
			break
		}
		u := n.SitesQueue[0]
		_, exist := n.Sites[u]
		if !exist {
			fmt.Println(i, "/", info.SearchNum, ":", u)
			site := GetSiteInfo(u)
			if site != nil {
				n.Sites[site.Url] = site
				for link := range site.Links {
					n.SitesQueue = append(n.SitesQueue, link)
				}
			}
		}
		n.SitesQueue = n.SitesQueue[1:]
	}
}
