.PHONY: docker
docker:
	@rm webook || true
	GOOS=linux GOARCH=arm go build -o webook .
	@docker rmi -f wulinlin/webook:v0.0.1
	@docker build -t wulinlin/webook:v0.0.1 .