c_m:
	migrate create -ext sql -dir db/migrations -seq $(name)

p_up:
	docker-compose up -d

p_down:
	docker-compose down

db_up:
	docker exec -it sip_pad_postgres createdb --username=root --owner=root sip_pad_db

db_down:
	docker exec -it sip_pad_postgres dropdb --username=root sip_pad_db

m_up:
	migrate -path db/migrations -database "postgres://postgres:secret@192.168.1.204:5432/sip_pad_db?sslmode=disable" up

m_down:
	migrate -path db/migrations -database "postgres://postgres:secret@192.168.1.204:5432/sip_pad_db?sslmode=disable" down

sqlc:
	sqlc generate

start:
	CompileDaemon -build="go build -o app.exe main.go" -command="./app.exe"

test:
	go test -v -cover ./...

