ENV=local
COVERFILE=cover.out
BASE_PATH=/go/src/github.com/atulsinha007/uber


test:
	ENV=$(ENV) BASE_CONFIG_PATH=$(BASE_CONFIG_PATH) go test -count 1 -p 1 -covermode count \
	-coverpkg=./... -v ./... -coverprofile $(COVERFILE) && go tool cover -func=$(COVERFILE)


test-cover:
	go tool cover -html=$(COVERFILE)


docker-test:
	docker-compose -f ./build/docker-compose.yml up --build --exit-code-from uber
	#docker-compose -f ./build/docker-compose.yml down --volumes


docker-test-coverage:
	bash ./build/tests_coverage_check.sh


end-test:
	COMPOSE_DOCKER_CLI_BUILD=1 docker-compose -f ./build/docker-compose.yml down --volumes


migrate_up:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/uber?sslmode=disable" up

migrate_down:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/uber?sslmode=disable" down