cdmain:
	cd ./vcs_ipatser

goenv:
	export GO111MODULE=auto
	export GOROOT=/usr/local/go
	export GOTOOLDIR=/usr/local/go/pkg/tool/linux_amd64
	export GOPATH=$HOME/workspace
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOBIN
	export GOMOD=$GOPATH/src/github.com/ipatser/go.mod

gobuild:
	cd ./vcs_ipatser; \
	go build

goinstall:
	cd ./vcs_ipatser; \
	go install

gorun:
	cd ./vcs_ipatser; \
	./vcs_ipatser;

gotest:
	cd ./vcs_ipatser; \
	bash test.sh; \
	rm test.out;

gofmt:
	gofmt -w .

dcpipe:
	docker-compose down; \
	docker image prune -f; \
	docker container prune -f; \
	docker rmi ipatser/build; \
	docker rmi ipatser/deploy; \
	docker-compose up
