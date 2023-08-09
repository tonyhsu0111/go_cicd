#CreateName="2019/12/24"
#MAINTAINER="Tony Hsu"
help:
	@echo ""
	@echo ""
	@echo "-- Help Menu"
	@echo ""
	@echo "   1. make windows        - build windows"
	@echo "   2. make linux          - build linux"
	@echo "   4. make build          - build golang file"
	@echo "   5. make dockerbuild    - build docker default"
	@echo "   5. make dockerclean    - clean docker image"
	@echo "   6. make clean          - Clean All File"


FileName="gocicd"

.PHONY:windows
windows:
	@echo "Build Windows filename ${FileName}.exe"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${FileName}_win.exe main.go  


build:
	go build -v -a -o release/linux/amd64/${FileName}

dockerbuild:
	@echo "Build Docker"
	@docker-compose -f docker-compose.yml build --force-rm

rm:
	@echo "remove Docker"
	@docker-compose -f docker-compose.yml down


run:
	@echo "RUN Docker"
	@docker-compose -f docker-compose.yml down
	@docker-compose -f docker-compose.yml up -d

logs:
	@docker-compose logs --tail 25 -f 

goclean:
	@rm -rf release

dockerclean:
	@echo "Remove image Docker"
	@docker-compose -f docker-compose.yml  down --rmi all
	@docker system prune -f

test:
	go test -v .