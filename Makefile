SERVER_NAME = gwk

build:
	@go build -o $(SERVER_NAME)

clean:
	@rm $(SERVER_NAME)

mock:
	@go build -o $(SERVER_NAME)

	@./gwk -task 'mock' -param '{"model": 1}'
	
	@rm $(SERVER_NAME)