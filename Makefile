dev:
	go run main.go

build:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app -a -ldflags '-w -s' main.go
