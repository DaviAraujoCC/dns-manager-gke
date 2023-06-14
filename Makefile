IMG =? controller:latest

all: build

build:
	@CGO_ENABLED=0 go build -o bin/dns-manager-gke main.go

docker-build:
	@echo "Building Docker image..."
	docker build -t ${IMG} .

docker-push:
	@echo "Pushing Docker image..."
	docker push ${IMG}
