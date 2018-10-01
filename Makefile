version=$(shell cat version)

build:
	go build

docker: build
	docker build . -t craftypenguins/minimal-mail-sender:$(version)

push:
	docker push craftypenguins/minimal-mail-sender:$(version)
