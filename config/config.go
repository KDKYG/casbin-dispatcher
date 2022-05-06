package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	globalConfig *GlobalConfig
)

var (
	confPath string
)

func GetGlobalConfig() *GlobalConfig {
	return globalConfig
}

type GlobalConfig struct {
	ServerID      string `yaml:"serverID"`
	DataDir       string `yaml:"dataDir"`
	JoinAddress   string `yaml:"joinAddress"`
	ListenAddress string `yaml:"listenAddress"`
	ServerPort string `yaml:"serverPort"`
}

func Init() {
	flag.StringVar(&confPath, "config-file", "", "Usage:./out - config-file xx")
	//解析上面定义的标签
	flag.Parse()

	globalConfig = &GlobalConfig{}
	fileBuffer, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalln("config ReadFile error:" + confPath)
	}
	err = yaml.Unmarshal(fileBuffer, globalConfig)
	if err != nil {
		log.Fatalln("config Unmarshal failed")
	}
	log.Printf("parse config :%v", *globalConfig)
}
