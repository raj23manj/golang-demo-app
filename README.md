# golang-demo-app

# Dev :
 - https://github.com/gin-gonic/gin/issues/477
 - nodemon --exec go run main.go --signal SIGTERM

# Run Migrations
- using go migrate => https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
- create a file:
  * migrate create -ext sql -dir db/migrations -seq create_tenants_table (if -seq is not given then timestamp is used)
  * export POSTGRESQL_GOLANG_URL='postgres://postgres:password@localhost:5432/golang_demo?sslmode=disable'
  * migrate -database ${POSTGRESQL_GOLANG_URL} -path db/migrations up
  * migrate -database 'postgres://postgres:password@localhost:5432/golang_demo?sslmode=disable' -path db/migrations up
  * when becomes dirty => migrate -database ${POSTGRESQL_GOLANG_URL} -path db/migrations force 1
