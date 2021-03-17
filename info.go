package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type infoStruct struct {
	LoadPath        string `yaml:"load_path"`
	OutputDir       string `yaml:"output_dir"`
	SearchNum       int    `yaml:"search_num"`
	OutputImageSize int    `yaml:"output_image_size"`

	Url       string `yaml:"url"`
	UserAgent string `yaml:"user_agent"`

	UrlFilter struct {
		UrlRegex    string `yaml:"url_regex"`
		DomainRegex string `yaml:"domain_regex"`
		PathRegex   string `yaml:"path_regex"`
	}

	SiteSettings struct {
		TitleRegex string `yaml:"title_regex"`
	}
}

func (info *infoStruct) readConf() {
	confFile, err := ioutil.ReadFile("info.yml")
	if err != nil {
		log.Println("Err: info file load err, ", err)
	}
	err = yaml.Unmarshal(confFile, info)
	if err != nil {
		log.Fatalln("Fatal: unmarshal, ", err)
	}
}

var infoInstance *infoStruct

func GetInfo() *infoStruct {
	if infoInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if infoInstance == nil {
			infoInstance = &infoStruct{}
		}
		infoInstance.readConf()
	}
	return infoInstance
}
