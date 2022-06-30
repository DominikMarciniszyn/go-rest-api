VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      go-rest-api

all: install

install:
	go install -v

test:
	go test ./... -v

fmt:
	go fmt ./... -v

image:
	docker build -t go-rest-api .

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

.PHONY: install test fmt release
