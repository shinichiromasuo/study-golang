NAME      := start-golang-api
REGISTRY  := acht/$(NAME)
VERSION   := debian-stretch-v2

.PHONY: build tag push

build:
	docker build -t $(NAME):$(VERSION) $(PWD)

tag: build
	docker tag $(NAME):$(VERSION) $(REGISTRY):$(VERSION)

push: tag
	docker push $(REGISTRY):$(VERSION)

go-build:
	GOOS=linux GOARCH=amd64 go build -o golang-api-fcgi server.go
