### Docker build function
## $1 TAG

docker-build=docker build \
	--compress \
	--rm \
	$(foreach tag, $1, -t $(tag)) \
	.


docker-push:
	docker push 