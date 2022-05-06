package casbin

import (
	"github.com/KDKYG/casbin-dispatcher/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	hraftdispatcher "github.com/casbin/hraft-dispatcher"
	"log"
)

// PERM
const modelConfig = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
`

var (
	globalEnforcer *casbin.DistributedEnforcer
)

func GetEnforcer() *casbin.DistributedEnforcer {
	return globalEnforcer
}

func Init() {
	var err error
	m, _ := model.NewModelFromString(modelConfig)
	globalEnforcer, err = casbin.NewDistributedEnforcer(m)
	if err != nil {
		log.Fatalln("NewDistributedEnforcer error:", err)
	}
	// New a Dispatcher
	dispatcher, err := hraftdispatcher.NewHRaftDispatcher(&hraftdispatcher.Config{
		Enforcer:      globalEnforcer,
		JoinAddress:   config.GetGlobalConfig().JoinAddress,
		ListenAddress: config.GetGlobalConfig().ListenAddress,
		DataDir:       config.GetGlobalConfig().DataDir,
	})
	if err != nil {
		log.Fatal(err)
	}

	globalEnforcer.SetDispatcher(dispatcher)
}

