.PHONY: clean dev-server build-server

clean:
	rm -rf todo.db
build-sever:
	go build
dev-server:
	go run main.go
