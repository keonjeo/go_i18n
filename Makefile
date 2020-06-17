.PHONY : dep test

dep:
	@echo "installing dependence..."
	@go get -v && go mod tidy

test:
	@echo "testing..."
	@go test
