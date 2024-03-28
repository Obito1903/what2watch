.PHONY: run build run-docker

run:
	go run cmd/db/main.go

build:
	docker build -t db .

run-docker: build
	docker run -p 8123:8080 db
