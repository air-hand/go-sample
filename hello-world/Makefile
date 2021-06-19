CONTAINER_BUILD_TARGET=prod

.PHONY: all
all: build
	docker compose up;

.PHONY: build
build:
	CONTAINER_BUILD_TARGET=$(CONTAINER_BUILD_TARGET) docker compose build;

.PHONY: shell
shell: CONTAINER_BUILD_TARGET=builder
shell: build
	docker compose run --rm helloworld /bin/zsh;

.PHONY: stop
stop:
	docker compose stop;

.PHONY: clean
clean: stop
	docker compose rm -f;
