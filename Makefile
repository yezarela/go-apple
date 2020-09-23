
run: 
	@echo "\Running app"
	@go run main.go

test:
	@echo "\nRunning tests"
	@go test -short  ./...