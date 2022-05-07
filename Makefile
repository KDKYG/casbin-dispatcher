APP=casbin-dispatcher

clean:
	rm -f ./${APP}

build:
	go build -o ${APP} ./main.go

run-leader: build
	./${APP} --config-file ./node-leader-conf.yml

run-follower1: build
	./${APP} --config-file ./node-follower-conf1.yml

run-follower2: build
	./${APP} --config-file ./node-follower-conf2.yml

