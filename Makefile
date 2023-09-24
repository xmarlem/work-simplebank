
.DEFAULT_GOAL := help

.PHONY: createdb
createdb: ## create a new database in postgres
	createdb --username=root --owner=root work_simplebank

.PHONY: dropdb
dropdb: ## delete the database in postgres
	dropdb work_simplebank


.PHONY: migrate_up
migrate_up: ## migrate the database schema up
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/work_simplebank?sslmode=disable" -verbose up

.PHONY: migrate.down
migrate_down:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/work_simplebank?sslmode=disable" -verbose down

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: sqlc
sqlc: ## generate sqlc code
	sqlc generate

.PHONY: help
help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
