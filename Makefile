build:
	go build

server:
	go run main.go

.PHONY: clean

clean:
	rm -f main
