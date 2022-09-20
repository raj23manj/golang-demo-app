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
* To get the type of the variable
  fmt.Println("%T")
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
* Concurrency
  * wait group
    * Go runtime package 50:00 OS, arch, cps, goroutines
    * init() 3:58
    * package sync
  * MethodSets
    * https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver
    * 11:40, methodset with interfaces not working with pointers
    * https://play.golang.org/p/G3lEy-4Mc8  for  below example
    ```
      package main

      import (
        "fmt"
        "math"
      )

      type circle struct {
        radius float64
      }

      type shape interface {
        area() float64
      }

      func (c *circle) area() float64 { // func (c circle) // works
        return math.Pi * c.radius * c.radius
      }

      func info(s shape) {
        fmt.Println("area", s.area())
      }

      func main() {
        // c:= &circle{5} // works
        c := circle{5}
        //info(c) // https://gobyexample.com/interfaces how to use interface
        // info(&c) works with (c *circle)
        fmt.Println(c.area())
      }
    ```
    * https://gobyexample.com/interfaces how to use interfaces

  * Documentation
    * 4:17, Go encourages "Do not communicate by sharing memory, instead, share memory by comunnicating"
    * channels(values are passed around on channels) is the way to go. 4:43 => dataraces cannot occur by design as one go routine has access to the value at any given time. => https://go.dev/doc/effective_go#concurrency
  * Race Condition
    * time.sleep() / runtime.Gosched() (run anything you want), 6:00
  * Mutex(mutual exclusion lock)
  * Atomic
    * similar to mutex, but like a  wrapper

* Channels
  * Notes:
    * to make the main program wait for exection and keep blocking
       1) use wait groups
       2) use a recieve channle in the main thread
       3) use range(will wait until the channle is closed)
       4) use select
  * understanding channels
    * 4:00, pass a baton at a same time.
    * send and recieve blocks until there is a channel on the other side
    * buffered channels 8:00,
    * try to stay away from buffered channels, use unbuffered channels. 10:40
  * Directional channels
  * Using Channels
    * 3:34, syntax for using send and receive channels
  * Range
    * range will keep looping over a channel, until a channel is closed. 1:20
    * main program
      ```
        func main() {
          c := make(chan int)
          go foo(c)

          for v := range c {
            fmt.Println(v)
          }
        }

        func foo(c chan<- int) {
          for i:=0; i < 100; i++ {
            c <- 42
          }
          close(c) // explicitly close, for range to quit
        }
      ```
      * when using range with channels, and we dont close the channel then we get a deadlock error
  * Select
   * it is like a switch statement
   * always need a case to hanlde quit the channel
   * main program
   ```
    func main() {
      eve := make(chan int)
      odd := make(chan int)
      quit := make(chan int)

      go send(eve, odd, quit)

      receive(eve, odd, quit)
    }

    func receive(eve, odd, quit <-chan int) {
      for {
        select {
          case v := <- eve:
            fmt.Println("from the even:", v)
          case v := <- odd:
            fmt.Println("from the odd:", v)
          case v := <- quit:
            fmt.Println("quit:", v)
            // close(quit)
            return
        }
      }
    }

    func send(e, o, q, chan<- int){
      for i:=0; i < 100; i++ {
        if i%2 == 0 {
          e <- i
        } else {
          o <- i
        }
      }
      // to make sure we close the channels, but this will be garbage collected if it was within this function
      close(e) // sends a value 0
      close(o)  // sends a value 0
      q <- 0
    }
   ```
    * running the program with race flag 9:20

  * Comma ok idiom
    * 0 is a false in go, 4:14
    * main program
    ```
    func main() {
      c := make(chan int)
      go func() {
        c <- 42 // send 0 it will be 0 true
        close(c) // send 0 and false, 5:56
      }()

      v, ok := <-c
      fmt.Println(v, ok) // 42 true

      v, ok = <-c
      fmt.Println(v, ok) // 0 false

    }
    ```
  * Fan In
    * get computation done and fan in all the results to main thread
    * https://go.dev/play/p/_CyyXQBCHe
    * https://go.dev/play/p/buy30qw5MM => rob pike, return channles from functions
     ```
      -------
             \
      -------------  main thread
             /
      -------
    ```
  * Fan Out & throttling
    * 3:17, closure concept
    * https://go.dev/play/p/iU7Oee2nm7, launches n go routines
    * throttling go routines
    ```
          -----------
         /
    ------------------
         \
          -----------
    ```
  * Context
