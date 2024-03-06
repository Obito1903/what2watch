.PHONY: build run-docker

run:
	npm run dev

build:
	docker build -t front .

run-docker: build
	docker run -p 3000:3000 front
