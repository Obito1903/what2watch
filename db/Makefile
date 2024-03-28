.PHONY: run build run-docker

run:
	uvicorn main:app --reload

build:
	docker build -t api .

run-docker: build
	docker run -p 8000:8000 api
