fmt:
	go fmt ./...

tidy:
	go mod tidy

build-local: fmt tidy
	go build -o bin/rfe-local -tags="local"

build-production:
	GOOS=linux GOARCH=amd64 go build -o bin/rfe-production -tags="production"

run-local: build-local
	sleep 0.5
	./bin/rfe-local
