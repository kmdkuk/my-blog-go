dev:
	docker-compose down
	docker-compose up -d
log-app:
	docker-compose logs -f app
build:
	branch=$(git branch --contains)
	docker build -t myblog-go:${branch} .
lint:
	goimports -w main.go
	go vet ./...
	golint ./...
