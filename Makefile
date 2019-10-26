build-raspi:
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o=./HikakinTeikyo main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o=./HikakinTeikyo main.go

build-cloudfunctions:
	zip upload.zip info  main.go  models  twitter  util  videoinfo.txt  youtube go.*
