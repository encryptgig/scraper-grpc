# Expands to list this project's go packages, excluding the vendor folder

TARGET = encryptgig


build: builddir
	GO111MODULE=on go build -o build/$(TARGET) ./server

builddir:
	@if [ ! -d build ]; then mkdir build; fi

update:
	GO111MODULE=on go get -u ./...

vendor.update:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

clean:
	rm -rf build/*

fmt:
	go fmt $(PACKAGES)

image:
	docker build . -t encryptgig


