default:
	godep go build -o bin/calendar src/api/api.go
test:
	godep go test
