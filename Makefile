run:
	GIN_MODE=release go run cmd/basic-golang-staticfile-server/main.go

dev:
	GIN_MODE=debug go run cmd/basic-golang-staticfile-server/main.go

build:
	go build cmd/basic-golang-staticfile-server/main.go