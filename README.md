# golang-demo-app

# Dev :
 - `go mod download` to install deps
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

# Log rocket
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/


# Timelines

1) why pass *domain.user
  * section2, package organization, 17:25
2) Basic setup:
  * section 2, package organization
3) Error handling
  * section 2, 32:31
4) Testing unit/integration/Functional test
            85/10/5
  * Section 3, intro to testing
5) Don't use panic
  * section 3, unit tests, 17:53
6) Using var in the service for exposing
  * Testing in go, 8:01 (how to structure go artifacts and mocks)
  * assume struct is like a class, and methods defined on it and exposing them
  * use interface for mocks 19:31
  * using init() menthod for mock 28:02
  * making dynamic functions for every test
