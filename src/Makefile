
build:
	go build -o main.exe main.go


buildx:
	GOOS=darwin go build -o main.exe main.go

run:
	GODEBUG=cgocheck=0 go run main.go


install:
	go get ./...
