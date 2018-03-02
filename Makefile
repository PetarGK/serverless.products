build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/handlers/addProduct src/handlers/addProduct.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handlers/listProducts src/handlers/listProducts.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/handlers/deleteProduct src/handlers/deleteProduct.go