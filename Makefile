default: run

run:
	go run main.go gbot.go

start:
	docker-compose up -d

stop:
	docker-compose down -v

build:
	go build -o gbot main.go gbot.go

docker: docker-build docker-push

docker-build:
	docker build -t bando/gbot .

docker-push:
	docker push bando/gbot
