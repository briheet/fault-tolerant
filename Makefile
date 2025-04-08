run: build
	@./cmd/bin

CMD: 
	@mkdir -p cmd

build: CMD
	@go build -o cmd/bin
