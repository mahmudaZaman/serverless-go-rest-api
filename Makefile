build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/api src/main/go/main.go
run:
	go run handler/hello/main.go