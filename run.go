package main

import "fmt"

func run(n *net) {
	info := GetInfo()

	for i := 1; i <= info.SearchNum; i++ {
		if len(n.RemainSites) <= 0 {
			break
		}
		u := n.RemainSites[0]
		_, exist := n.Sites[u]
		if !exist {
			fmt.Println(i, "/", info.SearchNum, ":", u)
			site := GetSiteInfo(u)
			if site != nil {
				n.Sites[site.Url] = site
				for link := range site.Links {
					n.RemainSites = append(n.RemainSites, link)
				}
			}
		}
		n.RemainSites = n.RemainSites[1:]
	}
}
