docker:
	docker pull mongo:3.4.0
	docker run -d -p 27017:27017 --name mongodb mongo:3.4.0

stop-docker:
	docker stop mongodb

build:
	go build -o main cmd/http/main.go

run:
	./main