
build:
	go build -o main.exe main.go


buildx:
	GOOS=darwin go build -o main.exe main.go

run:
	GODEBUG=cgocheck=0 go run main.go


install:
	go get ./...


#install npm nodemon
#sudo setcap cap_net_raw+eip $(eval readlink -f which node)


	