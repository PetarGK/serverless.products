build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/functions/createProduct/main functions/createProduct/main.go