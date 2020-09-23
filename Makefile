
run: 
	@echo "\nServing app"
	@go run main.go

test:
	@echo "\nRunning tests"
	@go test -short  ./...