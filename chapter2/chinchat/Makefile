.PHONY: build
build:
	go build

.PHONY: docker-image
docker-image:
	docker build -t chinchat .

.PHONY: clean
clean:
	rm -f chinchat
	docker image prune -f

start-pg:
	docker run -d --rm --name pg13 -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_HOST_AUTH_METHOD=trust postgres:13-buster

# start-psql:
# 	docker run -it --rm postgres:13-buster psql -h 172.17.0.4 -U postgres
#	\conninfo
#	\l	# show databases
