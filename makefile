c_m:
	# created a new migrations
	migrate create -ext sql -dir db/migrations -seq $(name)

p_up:
	# postgress up - create postgress server
	docker-compose up -d

p_down:
	# postgres down - delete postgres server
	docker-compose down

db_up:
	docker exec -it sip_pad_postgres createdb --username=root --owner=root sip_pad_db

db_down:
	docker exec -it sip_pad_postgres dropdb --username=root sip_pad_db

m_up:
	# run migrate up
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/sip_pad_db?sslmode=disable" up

m_down:
	# run migrate down
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/sip_pad_db?sslmode=disable" down

sqlc:
	sqlc generate

start:
	CompileDaemon -command="./finbest_backend"

test:
	go test -v -cover ./...

