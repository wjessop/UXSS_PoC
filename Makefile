all: assets
	goop go build -o bin/uxss

assets:
	go-bindata data/...

run: all
	bin/uxss -addr localhost:8080

clean:
	rm -f bin/*

deps:
	goop install

update_deps:
	goop update

releases: assets
	goop go build -o bin/uxss.mac
	GOOS=linux goop go build -o bin/uxss.linux
	GOOS=linux goop go build -o bin/uxss.exe

setup:
	go get -u github.com/jteeuwen/go-bindata/go-bindata...
	go get -u github.com/nitrous-io/goop
