include app.env
export $(shell sed 's/=.*//' app.env)

export GO111MODULE=on

prod_docker_compose_file=./docker/shops/docker-compose.yml

default: build test

# checks the code for problems
.PHONY: vet
vet:
	go vet

# starts the app
run: vet
	@echo "=============vetting the code============="
	go run main.go wire_gen.go server.go

.PHONY: mocks
mocks: clean-mocks
	@echo "=============generating mocks============="
	grep -rl --exclude "./vendor/*" --include "*.go" "interface {" . | while read -r file ; do mockgen --source=$$file --destination mocks/$$file ; done

# builds the executable
build:
	go build -o bin/shops main.go

# builds the docker image
build-docker:
	@echo "=============building shops============="
	docker build -f docker/shops/Dockerfile -t shops . ;

# sets up the hacker suite docker network
setup-network:
	@echo "=============setting up the hacker suite network============="
	docker network create --driver bridge hacker_suite || echo "This is most likely fine, it just means that the network has already been created"

# starts the app and MongoDB in docker containers
up: build-docker setup-network
	@echo "=============starting shops============="
	docker-compose -f $(prod_docker_compose_file) up -d

# starts the app and MongoDB in docker containers for dev environment
up-dev: export ENVIRONMENT=dev
up-dev: export PORT=8010
up-dev: export MONGO_HOST=127.0.0.1:8012
up-dev: vet setup-network
	@echo "=============starting shops (dev)============="
	refresh run

# prints the logs from all containers
logs:
	docker-compose -f $(prod_docker_compose_file) logs -f

# prints the logs only from the go app
logs-app:
	docker-compose -f $(prod_docker_compose_file) logs -f shops

# prints the logs only from the database
logs-db:
	docker-compose -f $(prod_docker_compose_file) logs -f mongo

# shuts down the containers
down:
	docker-compose -f $(prod_docker_compose_file) down

# cleans up unused images, networks and containers
# WARNING: this will delete ALL docker images on the system
#          that are not being used
clean: down
	@echo "=============cleaning up============="
	rm -f shops
	docker container prune -f
	docker network prune -f
	docker system prune -f
	docker volume prune -f

clean-mocks:
	rm -rf ./mocks
