#binance trader lite
binance-build:
	docker build --no-cache -t ${DOCKER_REPOSITORY_URL}/binance-fairy:v1 -f ./binance-fairy/Dockerfile ./binance-fairy
binance-start:
	docker-compose -f common/composers/binance-fairy.yml up -d
binance-stop:
	docker-compose -f common/composers/binance-fairy.yml down


# Common Prepare
development-build:
	docker build -t ${DOCKER_REPOSITORY_URL}/go_development:v1 - < ./common/build/Dockerfile