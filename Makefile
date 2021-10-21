install:
	go install

test:
	go test ./...

lint:
	staticcheck ./...