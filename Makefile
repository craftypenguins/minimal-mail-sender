build:
	go build

docker: build
	docker build . -t minimal-mail-sender:latest
