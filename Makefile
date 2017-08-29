.PHONY: deps build

deps:
	glide i
	rm -rf ./vendor/github.com/docker/docker/vendor
	rm -rf ./vendor/github.com/docker/distribution/vendor

build: deps
	go build -ldflags "-s -w" -a -tags netgo -installsuffix netgo

rpm:
	./make-rpm
