tidy:
	go mod tidy

server:
	go run main.go

build-image:
	docker build -t authservice-image -f Dockerfile .

run-container:
	docker run --net=host --rm -d --name authservice authservice-image

.PHONY: tidy server