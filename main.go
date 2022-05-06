package main

import (
	"github.com/KDKYG/casbin-dispatcher/casbin"
	"github.com/KDKYG/casbin-dispatcher/config"
	"github.com/KDKYG/casbin-dispatcher/router"
)

func main(){
	config.Init()

	casbin.Init()

	router.InitRouter()
}
