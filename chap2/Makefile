GO     := $(shell which go)

BIN := golang-api
APP := github.com/$(GITHUB_USER)/study-golang/chap2

CODES     := $(shell find . -type f -name '*.go' | grep -v "^./vendor")
PACKAGES  := $(shell find . -maxdepth 1 -type d | grep -Ev "^(.|./\..*|./data|./vendor|./kubernetes|./bin)$$" | sed -E 's/$$/\/.../g' | xargs go list)
LDFLAGS   := -ldflags="-X \"main.buildNum=0\" -X \"main.buildAt=$(shell date +%FT%T%z)\" -X \"main.branch=$(BRANCH)\" -X \"main.revision=$(REVISION)\""
BIN_DIST  := $(BIN)-$(shell date '+%Y%m%d-%H%M')-$(REVISION)

POD_NAME := $(shell kubectl get po -o=jsonpath='{.items[?(@.metadata.labels.app=="golang-api")].metadata.name}')

VENDOR_MD5 := $(shell if [[ -f vendor/lock.md5 ]]; then cat vendor/lock.md5; fi)
GOPKG_MD5 :=
ifneq ($(shell type md5 2> /dev/null),)
GOPKG_MD5 :=  $(shell md5 -q Gopkg.lock)
endif
ifneq ($(shell type md5sum 2> /dev/null),)
GOPKG_MD5 :=  $(shell md5sum Gopkg.lock | sed -E 's/ .*//g')
endif


.PHONY:	build vendor clean


vendor: dep-install
ifneq ($(VENDOR_MD5),$(GOPKG_MD5))
	dep ensure
ifneq ($(shell type md5 2> /dev/null),)
	md5 -q Gopkg.lock >| vendor/lock.md5
else ifneq ($(shell type md5sum 2> /dev/null),)
	md5sum Gopkg.lock | sed -E 's/ .*//g' >| vendor/lock.md5
else
	@echo vendor/lock.md5 was not created 1>&2
endif
else
	@echo vendor/ is already up-to-date
endif


clean:
	rm -f $(BIN)
	rm -rf vendor
	rm -rf $(GOPATH)/pkg/linux_amd64/$(APP)/*
	rm -rf $(GOPATH)/pkg/darwin_amd64/$(APP)/*
	dep remove


dep-install:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif


dep-ensure: dep-install
	dep ensure


.PHONY: ls-packages
ls-packages:
	@echo $(PACKAGES)


build: vendor
	GOOS=linux GOARCH=amd64 $(GO) build -i -gcflags "-N -l" -o $(BIN) ./*.go


.PHONY: deploy
deploy:
	kubectl cp $(GOPATH)/src/$(APP)/$(BIN) $(POD_NAME):/usr/local/bin/golang-api-fcgi -c golang-api
	kubectl exec $(POD_NAME) -c golang-api supervisorctl restart golang-api-fcgi


.PHONY: local
local: build deploy
