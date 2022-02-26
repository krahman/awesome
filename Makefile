.PHONY: test
test:
	go test ./...

.PHONY: dockerBuild
dockerBuild:
	docker build -t docker-go-awesome:multistage -f build/Dockerfile .

.PHONY: dockerRun
dockerRun:
	docker stop awesome || true && docker rm awesome || true
	docker run --name awesome -d -p 38000:38000 docker-go-awesome:multistage

.PHONY: build
build: test dockerBuild dockerRun
	curl http://localhost:38000