build:
	docker build -f cmd/car-server/Dockerfile -t jdowni000/temp:v2 .

push:
	docker push jdowni000/temp:v2
