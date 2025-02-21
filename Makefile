all: test
test:
	go test -v -cover ./internal/service/product | grep -v ./internal/service/product/product_mock.go -coverprofile=coverage.out

template: test
	go tool cover -html=coverage.out -o coverage.html

cover-full:
	go tool cover -func=coverage.out