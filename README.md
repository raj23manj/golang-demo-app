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

# Todo
 * coverage
# repo
federicoleon/golang-tutorial
https://github.com/federicoleon
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
7) gin.new => returns a gin instance without any middle attached,
  * Http framework 7:20
  * gin.Default has loggers & recovery middleware attached
8) How to fect query params http framework
  * 16:50
9) Marshal, takess a input interface and create a valid json string
  * 12:26 defining our domain structs
  * 20:49 unmarshalling
  * building nested structs 30:43
10) Building a rest client section4
  * Rest api calls 23:17
  * why to use pointers in return types of a function 30:37
  * read from response body ioutil.ReadAll(response.body) 36:13
  * use defer.Body.close() to close the body 41:47
  * Always build your own rest client it is easy to handle situations. 47:01
11) writing own custom mocks(mocking native apis)
  * 7:56(try writing mock like section 3 testing in go)
  * for body need a invalid closer , 24:46
  * for body response with a closer, ioutil.NopCloser() 29:05
12) Building a skeletal controller, service, error, provider(putting all together)
  * 8:59
  * 22:28 adding env/secrets
13) Testing: unit & integration tests
  * 5:34, initalizing things before testing using TestMain()
