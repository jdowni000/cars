build:
	docker build -f cmd/car-server/Dockerfile -t jdowni000/temp:v2 .

push:
	docker push jdowni000/temp:v2

buildkh:
	docker build -f cmd/car-checker/Dockerfile -t jdowni000/car-checker:v1.1.1 .

pushkh:
	docker push jdowni000/car-checker:v1.1.1
