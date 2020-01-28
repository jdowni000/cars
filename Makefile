build:
	docker build -f cmd/car-server/Dockerfile -t jdowni000/temp:testing .

push:
	docker push jdowni000/temp:testing
