build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/functions/create/main functions/create/main.go