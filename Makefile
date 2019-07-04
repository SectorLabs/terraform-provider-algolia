build:
	go build -o terraform-provider-algolia_v0.1

format:
	go fmt

run: build
	./terraform-provider-algolia_v0.1

all: build
