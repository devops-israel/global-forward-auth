clean:
	go clean -i ./...

deps:
	dep ensure

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o global-forward-auth .

build_osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o global-forward-auth .