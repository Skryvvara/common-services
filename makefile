.PHONY: deps
deps:
	cd src; go mod tidy

.PHONY: run
run:
	cd src; go run .

.PHONY: build
build: deps
	cd src; go build -o common-serivces