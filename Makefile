DATE = $(shell date +'%Y%m%d')
IMAGE = docker.io/nueks/simple-server

help: all
all:
	@echo "make [build | push]"

build:
	docker build -t $(IMAGE):$(DATE) .

push:
	docker tag $(IMAGE):$(DATE) $(IMAGE):latest
	docker push $(IMAGE):$(DATE)
	docker push $(IMAGE):latest
