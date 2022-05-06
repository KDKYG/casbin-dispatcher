package router

import (
	"github.com/KDKYG/casbin-dispatcher/casbin"
	"github.com/KDKYG/casbin-dispatcher/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	WatcherRouter *gin.Engine
)

func InitRouter() {
	WatcherRouter = gin.Default()
	WatcherRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	WatcherRouter.PUT("/policies", func(context *gin.Context) {
		var data interface{}
		context.ShouldBindJSON(&data)
		err := context.ShouldBindJSON(&data)
		if err != nil {
			http.Error(context.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = casbin.GetEnforcer().AddPolicies(interface2rules(data))
		if err != nil {
			http.Error(context.Writer, err.Error(), http.StatusServiceUnavailable)
			return
		}
	})
	WatcherRouter.DELETE("/policies",func(context *gin.Context) {
		var data interface{}
		context.ShouldBindJSON(&data)
		err := context.ShouldBindJSON(&data)
		if err != nil {
			http.Error(context.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = casbin.GetEnforcer().RemovePolicy(interface2rules(data))
		if err != nil {
			http.Error(context.Writer, err.Error(), http.StatusServiceUnavailable)
			return
		}
	})
	WatcherRouter.GET("/enforcer",func(context *gin.Context) {
		var data []interface{}
		err := context.ShouldBindJSON(&data)
		if err != nil {
			http.Error(context.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		ok,err := casbin.GetEnforcer().Enforce(data...)
		if ok || err == nil {
			context.JSON(http.StatusOK,"Authorized")
		}
		context.JSON(http.StatusBadRequest,"Unauthorized")
	})
	if err := WatcherRouter.Run(config.GetGlobalConfig().ServerPort); err != nil {
		log.Fatalln("Run error:", err)
	}
}

func interface2rules(i interface{}) [][]string {
	rules := make([][]string, 0)
	for _, item := range i.([]interface{}) {
		tmp := item.([]interface{})
		rule := make([]string, 0)
		for _, t := range tmp {
			rule = append(rule, t.(string))
		}
		rules = append(rules, rule)
	}
	return rules
}