BINARY_NAME=random-joke

clean:
	@echo "cleaning binary $(BINARY_NAME) (mac)" ;\
	rm $(BINARY_NAME)

build: clean
	@echo "building $(BINARY_NAME) (mac)" ;\
    go build -o $(BINARY_NAME) cmd/cli/main.go
build-rest-api: clean
	@echo "building $(BINARY_NAME) (mac)" ;\
    go build -o $(BINARY_NAME) cmd/rest-api/main.go

run:
	@echo "running (mac)" ;\
	go run cmd/cli/main.go
run-race:
	@echo "running (mac)" ;\
	go run cmd/cli/main.go --race

run-rest-api:
	@echo "running (mac)" ;\
	go run cmd/rest-api/main.go
run-race-rest-api:
	@echo "running (mac)" ;\
	go run cmd/rest-api/main.go --race

test:
	go test ./...
test-race:
	go test -race -short ./...

