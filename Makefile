.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./notify/notify

build:
	GOOS=linux GOARCH=amd64 go build -o notify/notify ./notify
