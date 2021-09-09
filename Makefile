BUILD_TARGET=prod

IS_IN_CONTAINER := $(shell sh -c 'test -f /.dockerenv && echo 0 || echo 1')

.PHONY: all
all: build

.PHONY: build
build:
	if [ $(IS_IN_CONTAINER) -eq 0 ]; then \
		cd src; \
		go build -o /go/bin/app; \
	elif [ "$(BUILD_TARGET)" = "builder" ]; then \
		docker compose -f docker-compose.yml -f .devcontainer/docker-compose.extend.yml build; \
	else \
		docker compose build; \
	fi;

.PHONY: up
up: build
	if [ $(IS_IN_CONTAINER) -eq 0 ]; then \
		/go/bin/app; \
	else \
		BUILD_TARGET=$(BUILD_TARGET) docker compose up; \
	fi;

.PHONY: test
test:
	if [ $(IS_IN_CONTAINER) -eq 1 ]; then \
		echo "test should be run in container."; \
		return; \
	fi; \
	find ./src -name "go.mod" | xargs -I{} readlink -e {} | xargs -I{} dirname {} | \
	xargs -I{} sh -c "cd {}; go test -test.v -cover"; \

.PHONY: shell
shell: export BUILD_TARGET=builder
shell: build
	docker compose run --rm web /bin/bash;

.PHONY: stop
stop:
	BUILD_TARGET=$(BUILD_TARGET) docker compose stop;

.PHONY: clean
clean: stop
	BUILD_TARGET=$(BUILD_TARGET) docker compose down --volumes --remove-orphans;
