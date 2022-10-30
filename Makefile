build:
	go build -v -a -o release/linux/amd64/gocicd

docker:
	docker build -t tonyhsu0111/gocicd .

test:
	go test -v .