version=$(shell cat version)

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker: build
	docker build . -t craftypenguins/minimal-mail-sender:$(version)

push:
	docker push craftypenguins/minimal-mail-sender:$(version)
