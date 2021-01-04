# private tasks. 
go-clean:
	rm -f ./bin/todo

go-build:
	go build -o ./bin/todo ./cmd/todo

# cmd interfaces.
build:
	$(MAKE) go-clean
	$(MAKE) go-build

clean:
	$(MAKE) go-clean
