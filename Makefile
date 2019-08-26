outputDir = ./out/kafka-protobuf-console

all: clean test build

clean:
	rm -rf ./out
	GO111MODULE=on go mod tidy -v

test:
	go test ./...

build:
	@echo "Building '${outputDir}'..."
	go mod tidy -v
	go build -o ${outputDir}
