init:
	echo "Initialized MySQL"
	docker-compose run --rm db sh -c "mysql -uroot -p -hdb < /sql/init.sql"
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
