outputDir = ./out/kafka-protobuf-console

all: clean test build

clean:
	rm -rf ./out

test:
	GO111MODULE=on
	go test ./...

build:
	GO111MODULE=on
	@echo "Building '${outputDir}'..."
	go mod tidy -v
	go build -o ${outputDir}

install:
	GO111MODULE=on
	go install