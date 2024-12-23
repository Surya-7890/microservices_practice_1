COMPOSE_DIR = deployment
COMPOSE_FILE = docker-compose.yml

.PHONY: dev
dev: build up

.PHONY: build
build:
	cd $(COMPOSE_DIR) && docker compose build

.PHONY: up
up: 
	cd $(COMPOSE_DIR) && docker compose up -d

.PHONY: stop
stop: 
	cd $(COMPOSE_DIR) && docker compose down

.PHONY: clean
clean:
	cd $(COMPOSE_DIR) && docker compose down --volumes --remove-orphans

.PHONY: buf-gen
buf-gen: buf-build
	buf generate
	go run move_gen_files.go

.PHONY: buf-build
buf-build:
	buf build