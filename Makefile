build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/handlers/addProduct src/handlers/addProduct.go