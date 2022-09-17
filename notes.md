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
  * Testing controller mocking context 21:51
  * Mocking headers 54:07
14) From linear to concurrent
  * implementation of concurrent api, 11:17 channels block, there has to be another go routine to accept or send
    if not will throw erreo all goroutines are assleep - deadlock!
  * buffered channels, 16:14, 21:00, if 3 is given to make(chan string, 3) 3 items will be buffered and for the fourth one
    it will wait untli some go routine recieves, else throws an error deadlock!



# links:
https://go.dev/ref/spec => The Go Programming Language Specification
https://pkg.go.dev/std => std packages

# Points for not leaking
* always check in the documentation or stack overflow if there is a close
  method for the function used to avoid leaks.
  ex: db connection, file, I/O, network, response body close, channels ...

# Todd Go
* Pointers
  * method sets
* Application
 * Json docementation
   * seeing the documentaions 1:59, golang.org & godoc.org
     with examples
   * how to read documentaion 11:25
   * Json Marshal & Unmarshal
   * writer interface
   * sort & sort custom
   * bcrypt
