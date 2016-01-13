all:
	$(GOPATH)/bin/goop go build -o bin/uxss

run: all
	bin/uxss -addr localhost:8080

clean:
	rm -f bin/*

deps:
	$(GOPATH)/bin/goop install

update_deps:
	$(GOPATH)/bin/goop update