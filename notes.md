# repo
federicoleon/golang-tutorial
https://github.com/federicoleon

#Todo
* https://go.dev/doc/faq
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
    * https://go.dev/doc/faq#closures_and_goroutines
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
    * 2:07 (https://go.dev/blog/context), In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using

* Error Handling
  * understanding
    * go does not have exceptions, 1:16
      https://go.dev/doc/faq#exceptions
      https://blog.golang.org/error-handling-and-go
  * checking errors
  * printing and logging
    You have a few options to choose from when it comes to printing out, or logging, an error message:
    * fmt.Println()  -> std o/p
    * log.Println()  -> std o/p or write to file
    * log.Fatalln()  -> os.exit()/shutting down
    * log.Panicln()
      ○ deferred functions run
      ○ can use “recover”
    * panic()
    * 5:24, use defered functions => defer foo()
  * Recover
    * https://go.dev/blog/defer-panic-and-recover
    * defer ensures that the function runs, 2:20
    * Recover is only usefulf insided deffered functions, 6:05
  * Errors with info

* Go doc, writting documentation
  * introduction
  * go doc
  * godoc
  * godoc.org
  * writting documentation


# Concurrency in GO: Deepak
  * Introduction
    * Processes & Threads
      - Process
        - Process is an instance of the running program
        - Process provides environment for program to execute
        - When the program is executed the os creates a process and allocates memory in the virtual address space
        - The virtual address space will contain code segment which is a complied machine code, there is a data region which contains global variables, Heap segment for dynamic memory allocation and stack is used for storing local variables of a function..
        ```
          ----------------
          |     stack |  |
          ------------v---
          |              |
          |              |
          ------------^---
          |      Heap |  |
          ----------------
          |      Data    |
          ----------------
          |      Code    |
          ----------------

        ```
      - Threads
        - Threads are smallest unit of execution that CPU accepts
        - Process has atleast one thread, that is the main thread
        - Process can have multiple threads
        - Threads share the same adress space
        ```
          ---------------
          |      Heap    |
          ----------------
          |      Data    |
          ----------------
          |      Code    |
          ----------------
          |              |
          |  t1      t2  |
          | stack     s  |
          | registers r  |
          | PC        pc |
          |              |
           ---------------
        ```
        - Threads run independent of each other
        - OS scheduler makes scheduling decisions at the thread level and not process level
        - Threads can run concurrently or parallel taking turns on individual core
        - Thread states
          - When process is created it is put in the ready Q ina a `runnable state`
          - Once the os scheduler schedules the thread, it is given a time slice
          - if time slice is expired then that thread is preempted and placed back to ready Q
          - If a thread gets blocked on I/O or event wait operation like read/write on disk, network operation, db operation, or waiting for an event from other processes then it is placed in the waiting Q until the I/O operation is completed. Once the I/O operation is completed the thread is placed back to ready Q.

        - Context switches
          - they are considered to be very expensive.
          - the cpu spends time in copying the context of the current executing thread into the memory and restoring the context of the next executing thread.
          - The context switiching will be a waste of time as it does switching instead of executing.
          - Context switching of the threads of a same process is cheap compared to context switching threads of different process
          - context-switching.png
          - If we have many processes per process and try to scale then it will lead to `C10K` problem
          - C10K Problem
            - Scheduler allocates a process a time slice for execution on CPU core
            - This cpu time slice is divided equally among threads
            - c10k.png
            - if time slice alloted for a process is 10ms by the scheduler
              - number of threads = 2, then thread time slice is 5ms
              - number of threads = 5, then thread time slice is 2ms
              - number of threads = 1000, then thread time slice is 10 nanoseconds, in this case the cpu does only context switching and does not execute.
            - To do a job, a thread needs minimum time slice of 2ms(milliseconds)
            - To complete one cycle of exection for number of threads, each thread has to wait n seconds for it execution
              ```
              | scheduler period | Number of threads | Thread time slice |
                      2s                  1000                2ms
                      20s                 10000               2ms
              ```
            - Bigger the time slice for the thread, slower is the application
          - Fixed Stack size:
            - The OS gives a fixed stack size of 8mb for a thread(on my machine)
            - The actual stack size depends on the hardware
            - If I have 8gb of memeory, in therory I can create only 1000 threads
               8GB/8mb = 1000
            - The fixed stack size limits the amount of threads that we create to the amount of menmory we have
