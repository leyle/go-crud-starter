VERSION_FILE=./VERSION
VERSION=`cat $(VERSION_FILE)`
COMMIT_HASH=`git rev-parse HEAD`
BRANCH=`git symbolic-ref --short -q HEAD`

SERVER_SRC="./cmd/apiserver/..."
SERVER_BIN="./build/release/api-server"

# https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#pushing-container-images
IMAGE_PATH="ghcr.io/leyle/go-crud-starter"

server:
	go clean && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_HASH) -X main.Branch=$(BRANCH)" -o $(SERVER_BIN) $(SERVER_SRC)

test:
	make server
	./build/release/api-server -c ./samples/api-config.yaml

image:
	echo $(VERSION)
	make server
	docker build -t $(IMAGE_PATH):$(VERSION) ./build

testimg:
	echo "test docker image"
	make image
	docker-compose -f docker-compose.sample.yaml down -v
	docker-compose -f docker-compose.sample.yaml up

pushimg:
	echo "push image to gitlab container registry"
	make image
	docker push $(IMAGE_PATH):$(VERSION)

clean:
	docker-compose -f docker-compose.sample.yaml down -v
	rm ./build/release/*
