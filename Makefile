TESTS ?= .*

build:
	go test -o example.test -c -race -v ./tests

test: build
	sudo -E ./example.test -test.v -test.timeout 10m -test.parallel=16 -test.run=$(TESTS)


bench:
	go test -benchmem -parallel=16 -bench Benchmark* ./tests -run=.*
