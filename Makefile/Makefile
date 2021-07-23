.PHONY: build

echo:
	echo gopath: $$GOPATH"\n"
	echo post_url:"http://www.ruanyifeng.com/blog/2015/02/make.html"\n

SERVICE := hello
MAIN := cmd/
REGISTRY := etcd
AUTHOR := $(shell git log --pretty=format:"%an"|head -n 1)
VERSION := $(shell git rev-list HEAD | head -1)
BUILD_INFO := $(shell git log --pretty=format:"%s" | head -1)
BUILD_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)
CUR_PWD := $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

LD_FLAGS='-X "$(SERVICE)/version.TAG=$(TAG)" -X "$(SERVICE)/version.VERSION=$(VERSION)" -X "$(SERVICE)/version.AUTHOR=$(AUTHOR)" -X "$(SERVICE)/version.BUILD_INFO=$(BUILD_INFO)" -X "$(SERVICE)/version.BUILD_DATE=$(BUILD_DATE)"'

VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default: build

build:
		go build -tags "$(REGISTRY)" -ldflags $(LD_FLAGS) -gcflags "-N" -i -o ./bin/$(SERVICE) ./$(MAIN)
race:
		go build -ldflags $(LD_FLAGS) -i -v -o ./bin/$(SERVICE) -race ./$(MAIN)
dev: build
		cp $(CUR_PWD)/conf/conf_dev.toml $(CUR_PWD)/conf/conf.toml && ./bin/$(SERVICE) -v=true
clean:
		rm bin/*
gofmt:
		echo "正在使用gofmt格式化文件..."
		gofmt -s -w ${GOFILES}
		echo "格式化完成"
govet:
		echo "正在进行静态检测..."
		go vet $(VETPACKAGES)

cover:
		@echo "代码覆盖率测试"
	    go test -coverpkg="./..."   -tags "$(REGISTRY)" -ldflags  $(LD_FLAGS)  -c -cover -covermode=atomic -o ./bin/$(SERVICE) ./$(MAIN)

