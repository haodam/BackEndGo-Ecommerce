db_docs:
	dbdocs build doc/db.dbml
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

postgres:
	docker run --name ecommerce -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123secret -d postgres:16-alpine
create_db:
	docker exec -it ecommerce createdb --username=root --owner=root ecommerce
drop_db:
	docker exec -it ecommerce dropdb ecommerce

migrate_up:
	migrate -path migration -database "postgresql://root:123secret@localhost:5432/ecommerce?sslmode=disable" -verbose up
migrate_down:
	migrate -path migration -database "postgresql://root:123secret@localhost:5432/ecommerce?sslmode=disable" -verbose down

.PHONY: db_docs db_schema create_db drop_db postgres migrate_up migrate_down