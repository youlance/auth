tidy:
	go mod tidy

server:
	go run main.go

.PHONY: tidy server