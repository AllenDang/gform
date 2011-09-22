all: clean
	make -C pkg/gform		install

test: clean
	make -C pkg/gform		test

clean:
	make -C pkg/gform		clean

format:
	gofmt -w .
