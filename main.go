package main

func main() {
	info := GetInfo()

	n := net{ RootUrl: info.Url }
	n.Init()
	n.Run(info.SearchNum)

	outputMarkdown(&n, info.OutputDir)

}

