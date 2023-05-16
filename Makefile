#CreateName="2019/12/24"
#MAINTAINER="Tony Hsu"
help:
	@echo ""
	@echo ""
	@echo "-- Help Menu"
	@echo ""
	@echo "   1. make windows        - build windows"
	@echo "   2. make linux          - build linux"
	@echo "   4. make build          - build default"
	@echo "   5. make clean          - Clean All File"


FileName="gocicd"

.PHONY:windows
windows:
	@echo "Build Windows filename ${FileName}.exe"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${FileName}_win.exe main.go  


build:
	go build -v -a -o release/linux/amd64/${FileName}

docker:
	docker build -t tonyhsu0111/${FileName} .

test:
	go test -v .