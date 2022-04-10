IS_IN_CONTAINER := $(shell sh -c 'test -f /.dockerenv && echo 0 || echo 1')

.PHONY: migrate
migrate:
	if [ $(IS_IN_CONTAINER) -eq 1 ]; then \
		echo "should be run in container."; \
		return; \
	fi; \
	go run *.go $(MIG_COMMAND);

.PHONY: up
up: MIG_COMMAND=up
up: migrate

.PHONY: down
down: MIG_COMMAND=down
down: migrate

.PHONY: force
force: MIG_COMMAND=force $(VER)
force: migrate
