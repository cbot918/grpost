DB_URL=postgres://postgres:12345@localhost:5432/grpost?sslmode=disable

run: build-ui
	go run .

watch:
	gowatch

build-ui:
	npm --prefix ui install
	npm --prefix ui run build 

# database

mig-init:
	migrate create -ext sql -dir db/migrations -seq $(ARG)

mig-up:
	migrate -path db/migrations -database $(DB_URL) up

mig-down:
	migrate -path db/migrations -database $(DB_URL) down

mig-down:

.PHONY: run build-ui watch
.SILENT: run build-ui watch

test:
	echo $(ARG)