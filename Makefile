.PHONY: test run-all-local-logs run-all-local stop-discount-service start-discount-service down-all-local

test:
	- docker build --progress=plain \
		--target=test \
		--file ./build/package/dockerfile-test .

run-all-local:
	docker-compose -f build/package/docker-compose.yml up --build

down-all-local:
	docker-compose -f build/package/docker-compose.yml down

stop-discount-service:
	docker-compose -f build/package/docker-compose.yml stop discount-service

start-discount-service:
	docker-compose -f build/package/docker-compose.yml start discount-service
