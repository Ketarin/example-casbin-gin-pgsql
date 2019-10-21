start: build	
	@echo "Starting API."	
	./myapp
build:
	@echo "Building package..."
	go build -o ./myapp ./app/main.go
#DOCKER TESTS
complete: stop rebuild run
stop:
	docker-compose stop
rebuild:
	docker-compose rm -f 
	docker-compose build
run:
	docker-compose up 
