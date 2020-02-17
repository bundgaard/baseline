### Docker build function
## $1 TAG

docker-build=docker build --compress --rm .


docker-push:
	docker push 