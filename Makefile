CONTAINER_BUILD_TARGET=prod

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
	CONTAINER_BUILD_TARGET=$(CONTAINER_BUILD_TARGET) docker compose build;

.PHONY: up
up: build
	docker compose up;

.PHONY: shell
shell: CONTAINER_BUILD_TARGET=builder
shell: build
	docker compose run --rm web /bin/zsh;

.PHONY: stop
stop:
	docker compose stop;

.PHONY: clean
clean: stop
	docker compose rm -f;
