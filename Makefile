BUILD_TARGET=prod

.PHONY: all
all:
	IS_IN_CONTAINER=$$(test -f /.dockerenv && echo 0 || echo 1); \
	if [ $$IS_IN_CONTAINER -eq 0 ]; then \
		go build -o /go/bin/app; \
	else \
		make build; \
	fi

.PHONY: build
build:
	if [ "$(BUILD_TARGET)" = "builder" ]; then \
		docker compose -f docker-compose.yml -f .devcontainer/docker-compose.extend.yml build; \
	else \
		docker compose build; \
	fi

.PHONY: up
up: build
	BUILD_TARGET=$(BUILD_TARGET) docker compose up;

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
