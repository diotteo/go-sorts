.PHONY: all
all: build

.PHONY: build
build:
	mkdir build && go build -o build/ src/*.go

.PHONY: clean
clean:
	rm -r build/

.PHONY: test
test:
	go test -v ./src
