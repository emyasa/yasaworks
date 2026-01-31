build-ssh:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yasaworks ./cmd/ssh

