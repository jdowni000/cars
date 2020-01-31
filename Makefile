build:
	docker build -f cmd/car-server/Dockerfile -t jdowni000/temp:v3 .

push:
	docker push jdowni000/temp:v3
