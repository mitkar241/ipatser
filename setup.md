# Project Setup
---

## Action before Git clone
---

```bash
cd $HOME
mkdir -p workspace/src/github.com
git clone https://github.com/raktimhalder241/ipatser.git
```

## Set following environment variables in current terminal
---

```bash
# go env
export GO111MODULE=auto
export GOROOT=/usr/local/go
export GOTOOLDIR=/usr/local/go/pkg/tool/linux_amd64
export GOPATH=$HOME/workspace
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
export GOMOD=$GOPATH/src/github.com/ipatser/go.mod
```

## Create go.mod and go.sum
---

```bash
go mod init ipatser
go get github.com/lib/pq
go get github.com/gorilla/mux
```

## Build main binary `vcs_ipatser`
---

```bash
cd vcs_ipatser/
go build
go install
ls $GOBIN
```
```
vcs_ipatser
```

## Build other packages (optional)
---

```bash
cd utils/
go build
go install
ls $GOPATH/pkg/linux_amd64/github.com/ipatser/
```
```
utils.a
```
