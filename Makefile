GO  := $(shell which go)
APP := $(GOPATH)/src/github.com/$(GITHUB_USER)/study-golang
POD_NAME := $(shell kubectl get po -o=jsonpath='{.items[?(@.metadata.labels.app=="golang-api")].metadata.name}')

.PHONY: kube-create dev-deploy start

kube-clean:
	sed -e "s#{{ PROJECT_DIR }}#$(APP)#g" kubernetes/golang-api.yaml | \
	kubectl delete \
		-f kubernetes/nginx-conf.yaml \
		-f -

kube-create:
	sed -e "s#{{ PROJECT_DIR }}#$(APP)#g" kubernetes/golang-api.yaml | \
	kubectl apply \
		-f kubernetes/nginx-conf.yaml \
		-f -


dev-deploy:
	kubectl cp $(APP)/kubernetes/containers/go-fcgi/golang-api-fcgi $(POD_NAME):/usr/local/bin/golang-api-fcgi -c golang-api
	kubectl exec $(POD_NAME) -c golang-api supervisorctl restart golang-api-fcgi

mysql-create:
	docker-compose up -d

start: kube-create dev-deploy mysql-create
