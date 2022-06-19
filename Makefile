run:
	GIN_MODE=release go run cmd/basic-golang-staticfile-server/main.go

dev:
	GIN_MODE=debug go run cmd/basic-golang-staticfile-server/main.go

test:
	APP_NAME=test-app GIN_MODE=debug go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

build:
	go build cmd/basic-golang-staticfile-server/main.go