package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	AppUid    string `yaml:"app_uid"`
	AppSecret string `yaml:"app_secret"`
	Cursus    string `yaml:"cursus"`
}

func parseConfig() config {
	conf := config{}

	content, err := ioutil.ReadFile(os.Getenv("HOME") + "/.intraRpc.conf")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
