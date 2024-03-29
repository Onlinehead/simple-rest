IMAGE=onlinehead/simple-rest
POSTGRES_HOST=127.0.0.1
VERSION=0.1
BUILD_TIME=$(shell date +%Y-%m-%d-%H-%M-%S)

postgres_up:
	docker run -d --rm --name=postgres -e POSTGRES_PASSWORD=zzz -e POSTGRES_USER=user -e POSTGRES_DB=simple_rest -p 5432:5432 postgres

postgres_down:
	docker stop postgres

image:
	docker build -t $(IMAGE):$(VERSION) --build-arg Version=${VERSION} .

push:
	docker push $(IMAGE):$(VERSION)

run_local:
	GO111MODULE=on go run server.go

run:
	docker run -it -e REST_POSTGRESHOST=${POSTGRES_HOST} -p 8081:8080 --rm $(IMAGE):$(VERSION)

build:
	GO111MODULE=on go build -o bin/simple-rest -ldflags "-X BuildTime=`date +%Y-%m-%d:%H:%M:%S`" *.go

# If you are using remote Docker backend, you cannot mount your local dir with a code, so you need to build a container first
tests:
	docker build -t simple-rest-tests:${BUILD_TIME} -f Dockerfile_tests .
	docker run -it --rm -e GO111MODULE=on simple-rest-tests:${BUILD_TIME} go test
	docker rmi simple-rest-tests:${BUILD_TIME}

tests_local:
	GO111MODULE=on go test ./...

tests_all: postgres_up tests postgres_down
