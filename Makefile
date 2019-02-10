build-raspi:
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o=./bin/HikakinTeikyo main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o=./bin/HikakinTeikyo main.go
