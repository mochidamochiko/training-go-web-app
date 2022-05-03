.PHONY: build
build:
	go build

.PHONY: docker-image
docker-image:
	docker build -t first_webapp .

.PHONY: clean
clean:
	rm -f first_webapp
	docker image prune
