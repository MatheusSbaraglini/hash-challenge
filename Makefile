.PHONY: test run-local

test:
	- docker build --progress=plain \
		--target=test \
		--file ./build/package/dockerfile-test .

run-local:
	docker-compose -f build/package/docker-compose.yml up --build