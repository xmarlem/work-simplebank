

.PHONY: createdb
createdb:
	createdb --username=root --owner=root work_simplebank


.PHONY: dropdb
dropdb:
	dropdb work_simplebank


.PHONY: migrate.up
migrate.up:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/work_simplebank?sslmode=disable" -verbose up

.PHONY: migrate.down
migrate.down:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/work_simplebank?sslmode=disable" -verbose down
