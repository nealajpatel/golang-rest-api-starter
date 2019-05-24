NAME=golang-rest-api-starter
VERSION=0.0.1

.PHONY: build
build:
	@echo Building from source....
	@go build -o ./build/$(NAME)

.PHONY: run
run: build
	@echo Starting your app using dev configs....
	@./build/$(NAME) -e dev

.PHONY: build-prod
build-prod:
	@echo Building from source....
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

.PHONY: run-prod
run-prod:
	@echo Starting app using prod configs....
	@$CI_PROJECT_DIR/golang-rest-api-starter-binary -e dev

.PHONY: clean
clean:
	@echo Removing build file....
	@rm -f ./build/$(NAME)

.PHONY: test
test:
	@go test -v ./tests/*
