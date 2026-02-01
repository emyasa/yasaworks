run-cli:
	go run ./cmd/cli

build-ssh:
	go build -o yasaworks-mac ./cmd/ssh
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yasaworks ./cmd/ssh

