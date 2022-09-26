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
  * git clone https://github.com/andcloudio/go-concurrency-exercises.git
  * go-concurrency-exercises-ind
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

        - What are the limitations of thread ?
          * Fixed stack size
          * C10K problem, as we scale up the number of threads, scheduler cycle increases and application can become less responsive.

    * why concurrency is hard
      * shared memory
        - Threads communicate between each other by sharing memory.
        - Threads share the heap and data region of the process
        - sharing the memory between threads creates a lot of complexity with concurrently executing threads.
        - Cocnurrent access to shared memory by two or more threads can lead to `Data Race` and the outcome can be Un-deterministic.
        - example, 1:40
      * Memory Access Synchronization
        - we need to guard the access to shared memory so that a thread gets exclusive access at a time.(mutex locks / mutually exclusive locks)
        - Locking reduces parallelism. Locks force to execute threads sequentially.
        - inappropriate use of locks leads to deadlock.
        - deadlock.png

  * Goroutines
    * Goroutines
      - There are limitation with threads, with the actual number of threads we can create is limited, sharing of memory leads to lot of complexity, with concurrently executing threads.
      - Communicating Sequential Processes(CSP) Tony Hoare
        - Each process is built for sequential execution, every process has a local state. The process operates on that local state
        - Data is communicated between processes withou sharing memory, but we send a copy of the data to the other processes. Since there is no sharing of memory there will be no race condition or deadlocks
        - Scale easily, as each process run independently. If the computation is taking more time, we can add more processes of the same type.
      - Go's Concurrency Tool Set
        - goroutines
        - channels (communication)
        - select (multiplex the channels)
        - sync package(classical sync tools like mutex)
      - `Goroutines`
        - they are user space threads(like green threads) managed by go runtime
        - Go runtime is part of the executable that is built in
        - Goroutines are extremely lightweight.
        - Goroutines starts with `2KB` of stack, which grows and shrinks as required. Where as threads need 8mb(based on hardware)
        - Low CPU overhead - three instructions per function call.
          - the amount of cpu instructions to create a goroutine is very less
        - Can create hundreds of thousands of goroutines in the same address space.
        - Channels are used for communication of data between goroutines. Sharing of memory can be avoided.
        - Context switching between goroutines is much cheaper that thread context switching, as go routines have less state to store.
        - Go runtime can be more selective in what data is persisted for retrieval, how it is persisted, and when the persisiting needs to occur.
        - Go Runtime creates OS threads
        - Goroutines runs in the context os OS thread.
        - Many goroutines can execute commands in the context of Single OS thread.
        - The OS scheduler schedules the os threads, and the goruntime schedules the goroutines on the os thread.
        - os-go-runtime-scheduler.png

    * Exercise-Hello
      * use of sleep
        time.Sleep(1 * time.Millisecond)
      * goroutine function call
      * goroutine with anonymous function
      * goroutine with function value call
      * wait for goroutines to end
      * https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/01-goroutines/01-hello/main.go

    * Exercise-ClientServer
     * https://github.com/andcloudio/go-concurrency-exercises/tree/master/01-exercise-solution/01-goroutines/02-client-server

    * Waitgroups:
      - different possible outcome of the prog:
         ```
           func main() {
             var data int

             go func() {
               data++
             }()

             if data == 0 {
               fmt.Println("value %v", data)
             }
           }

           o/p(any of below):
           1) nothing is printed. If child Goroutine finishes before main Go routine it increments the value
           2) the value is 0, if main go routine executes before child Go routine
           3) value is 1, ig child goroutine gets scheduled in between data === 0 and next line of execution
         ```
      - Wait Group
        ```
          var wg sync.WaitGroup
          wg.Add(1) // indicates the number of go routines we are creating

          go func() {
            defer wg.Done() // need to call same number of times as we add in main go routine.
          }

          wg.Wait()
        ```
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/01-goroutines/03-join/main.go

    * Goroutine & Closure
      - Goroutines execute within the same address space they are created in
      - They can directly modify varaibles in the enclosing lexical block
      ```
        func inc() {
          var i int

          go func() {
            i++
            fmt.Println(i)
          }()

          return
        }
      ```
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/01-goroutines/05-closure/main.go
        - closure function, nested functions
        - 1:25, talks about varaible going out of scope once function returns, but the variable `i` is moved to heap due to closure
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/01-goroutines/06-closure/main.go
        - closure example why we need to pass values to function instead of referencing them from nested function using closure
        - By the time go routines got a chance to execute the main goroutine runs the loop and waits so i is 4
        - To avoid this, pass the value to the function
        - Go routines work on the current value that they have, to make this work properly pass the value

  * Go Scheduler
    * DeepDive - Go Scheduler
      - it part of the Go runtime, and runtime is part of the executable of the application
      - also know M:N scheduler
      - It runs in the user space
      - Go scheduler used OS threads to schedule goroutines for execution
      - Go routines runs in the context of OS threads.
      - Go Runtime create a number of worker OS threads, equal to `GOMAXPROCS`
      - `GOMAXPROCS` default value is the number of processors on the machine
      - Go Scheduler distributes runnable goroutines over multiple worker OS threads.
      - At any time, N Goroutines could be scheduled on M OS threads that runs on at most GOMAXPROCS number of processors.
      - As of GO 1.14 Go scheduler implements asynchronous preemption, before that it was co-operative premption.
      - asynchronous preemption prevents long running Goroutines from hogging onto CPU, that could block other GoRoutines
      - The asynchronous preemption is triggered based on a time slice. When a goroutine is running for `more that 10ms(milli)`, Go will try to preempt.
      - Goroutine States similar to threads, 2:49
        - When Goroutinr is created, it will be in runnable state
        - When executing, it goes to executing state when the goroutine is scheduled on a thread. If not completed with in the time slice, it is preemted and placed back into runnable state
        - If the Go routine is wating for blocked on a channel, blocked on a sys call, waiting for a mutex lock, I/O or event wait to be completed, it is placed under waiting state. Once the waiting operation is complete it is moved back to runnable state.
        - Goroutines-states.png
      - Go Runtime architecture 3:21
        - For a Cpu core, go runtime creates a Thread M
        - For a thread M, go runtime creates a Logical processor(with stack, head, data regions) P
          - The Logical processor holds the context for scheduling, which can be seen as a local scheduler running on a thread.
          - Each processor P has a local run queue(LRQ), where runnable goroutines are queued.
        - G represents a goroutine running on the OS thread
        - There is a Global Run Queue(GRQ), once LRQ is exhausted, the processor will pull go routines from GRQ.
        - New Goroutines are added to the end of the GRQ
        - scheduler-arch.png

    * DeepDive - GO Scheduler - Context switching due to Synchronous system call
      - what happens when a synchronous system call is made? like read/write to a file with sync flag set.
        - synchronous system calls wait(block) for I/O operations to complete.
        - when this happens, OS thread is moved out of the CPU to waiting queue for I/O operations to complete.
        - The scheduler won't be able to schedule any other Go routines on this thread.
        - The the OS runtime creates a new thread(or take a thread from thread pool) and moves all the LRQ of the blocked thread and attaches it to the new thread and attaches the new thread to the core.
        - Once the blocked thread goroutine finishes its sync operation, it will be move to one of the LRQ of the running thread(processor).
        - The blocked thread detached, will be put to sleep and placed in the thread pool. It will be utilized in future for same scenario.
        - Synchronous system calls reduces parallelism.
        - sync-systemcalls.png

    * DeepDive - Go Scheduler - Context switching due to Asynchronous system(network system call, http system call) call
      - what happens in general when asynchronous system calls are made?(scenarios below)
        - Asynchronous system call happens when the file descriptor(this is used for doing network I/O operations) is set to non-blocking mode.
        - If the file descriptor is not ready(if the socket buffer is empty for read operation or the socket buffer is full for write operation), system call does not block, but returns an error. If this happens, the application has to retry in a later point in time.
        - Asynchronous I/O increases the application complexity.
        - The application has to create a event loop and setup callbacks or it has to maintain a table mapping the file descriptor and the function pointer.(Go routines have these) It has to keep track of how much data was read last time or how much data was returned last time. All this add up to complexity of the application. If not implemented properly, it makes all in efficient.
      - How does Go handle the above scenarios?
        - Netpoller
          - Go uses netpoller, it is an abstraction built in syscall package
          - syscall package uses netpoller to conver asynchronous system calls to blocking system call.
          - When a goroutine makes asynchronous system call and the file descriptor is not ready, then the Go Scheduler uses netpoller OS thread to park that goroutine.
          - netpoller uses interface provided by OS to do polling on the file descriptor. It has the blocked goroutines in its data structure.
            - the netpoller uses
              * kqueue (Mac OS) to do polling on the file descriptor
              * epoll(Linux)
              * iocp(windows)
          - netpoller gets notification from OS, when the file descriptor is ready for I/O operations.
          - netpoller notifies goroutine to retry I/O operations, if there are any in its data structure.
          - To retry I/O operations, the Go routine is made runnable again and is moved from netpoller to one of the running processors LRQ
          - This way the complexity of managing asynchronous system call is moved from Application to Go Runtime, which manages it efficiently.
          - This way of using the netpoller, no extra threads are created for execution, but uses existing netpoller thread to get asynchronous call executed.
          - netpoller.png

    * DeepDive - Go Scheduler - Work Stealing
      - Work stealing helps to balance the goroutines acrss all the logical processors
      - work gets distributed and gets done more efficiently
      - work stealing rule
        - if there is no goroutines in LRQ
          - try to steal from other logical processors.
          - if not found, check the global runnable queue for G
          - if not found check the net poller
          - this repeats as cycle.
        - The logical processor randomly pick from other logical processors and steals half of it's G

  * Channels(hchan struct represents a channel)
    - Channels are used to communicate data between Goroutines
    - Channels help in synchronization of the execution of the go routines
    - One Goroutine can let know another Goroutine in what stage they are in and synchronize their execution
    - channels are typed.
    - channels(un-buffered) block the execution of the go routines, the sender go routine will block until there is a reciever goroutine ready and vice-versa.
    - channels are thread-safe and can be used safely to send and recieve values concurrently, by multiple goroutines.
    -  Default value of channels delared is nil
       ```
        var ch chan T
        ch = make(chan T)

        or

        ch := make(chan T)
       ```
    - Pointer operator is used for sending and receiving the value from the channel.
    - The arrow indicated the direction of data flow
      ```
        ch <- v // send data
        v = <- ch // receive data
        ```
    - is is the responsibilty of the channel to make the go routine runnable again once it has data.
    - to close the channel `close(ch)`. This is used for the sender goroutine to indicate the reciever that the sender has no more values to send.
    - once close the channel `close(ch)`, the reciever will be unblocked and computation will continue normally.
    - Recieve channel gets two values
      - `v,ok = <- ch`
      -  v is the value from the channel passed, ok = true, value generated by a write
      - ok = false, value generated by a close.
    - Excersice channels: https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/02-channel/01-channel/main.go
    - Range and Bufferef channels:
      - iterate over the values recevied from a channel
      - Loop automatically breaks, when a channel is closed
      - range does not return the second boolean value
        ```
            for value := range ch {

            }
        ```
      - Unbuffered channels
         - normal declared channels are unbuffered anf they block.
      - Buffered channel
        - There is a buffer specified between the sender and the receiver
        - There is a capacity mentioned, which indicates the number of elements that can be sent over a channel withou the `reciever` being ready to recieve the values.
        - The sender can keep sending the values without blocking, till the buffer gets full. Once the buffer gets filled the channel will block.
        - The receiver can keep receiving the values until the buffer gets empty. When the buffer gets empty the receiver will block.
        - Buffered channels are in memory `FIFO` queues, the element sent first will be received first.
    - Range ecxersice:
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/02-channel/02-channel/main.go
    - Bufferd channels excersice
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/02-channel/03-channel/main.go
    - Channel direction
      - When using channels as functional parameters, you can specify if a channel is menat to only send or receive values.
      - This specifically imcreases the type-safety of the program
      ```
        func pong(in <-chan string, out <-chan string) {}
      ```
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/02-channel/04-channel/main.go
    - Channel ownership
      - Default values foa channels whe declaredL nil
        ```
          var ch chan interface{}
        ```
      - We need use the `make` function to allocate memory
      - If we use channel without allocating memory, reading/writing to a nil channel will block forever.
        ```
          var ch chan interface{}
          <- ch
          ch <- struct{}{}
        ```
      - Closing on a nil channel will panic
        ```
          var ch chan interface{}
          close(ch)
        ```
      - Ensure the channels are initalized first with the `make` function
      - How to use the channels is important, to aboud dead locks and panics by following GO idioms
        - Best practises
          - Owner of a channel is a goroutine that instantiaites, writes and closes a channel.
          - The other goroutines that utlize the created channel will only have a read-only view into the channel.
      - Ownership of a channel avoids
        - Deadlocking by writing to a nil channel, leads to panic
        - closing a nil channel, leads to panic
        - writing to a closed channel, leads to panic
        - closing a channel more than once, leads to panic
      - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/02-channel/05-channel/main.go

  * Deep Dive Channel
    - To create a channel we use `ch := make(chan int, 3)`, internally channels are represented by `hchan` struct.
    - hchan-struct.png
    - The hchan structhas a `mutex lock` field, any goroutine doing a channel operation must first aquire a mutex lock first.
    - It has `buf` field which stores actual data. This is used only for the buffered channels.
    - `dataqsiz` field is the  size of the buffer.
    - `qcount` indicates the total data in the queue.
    - `sendx and recvx` indicates the current index of the buffer, form where it can send data and receive data.
    - `sendq & recvq` are used to store the blocked goroutines.
    - `wait q` is a struct represented by the `sudog` struct.
    - sudog-struct.png
       - `g`, represents the pointer to the goroutine
       - `elem` represents the ponter to memeory which contains the value to be sent or to which the received value will be returned to.
    - `ch := make(chan int,3)`
      - when make is called, hchan struct is allocated in the heap
      - make() returns a pointer to it
      - Since `ch` is a pointer, it can be between functions for send and receive.
    - channel-allocation.png

  * Send and Receive on buffered channels

    ```
      ch := make(chan int, 3)
      // G1 - goroutine
      func G1(ch chan<- int) {
        for _, v := range []int{1,2,3,4} {
          ch <- v
        }
      }

      // G2 - goroutine
      func G2(ch chan<- int) {
        for v := range ch {
          fmt.Println(v)
        }
      }
    ```
    - init-send-channel.png
    - working
      - G1 aquires the lock first
      - G1 enques the value on to the queue(it is a memory copy, element is copied into the buffer)
      - G1 increments the value of sendx to 1
      - G1 releases the lock
      - G1 go aheads with it other computation
      - G2 comes, aquires the lock first
      - G2 deques the value from the buffer queue and copies it to variable `v`
      - G2 increments the value of recvx
      - G2 releases the lock
      - G2 go aheads with it other computation
      - Points to note:
        - There is no memory sharing between Goroutines
        - Goroutines copy elements into and from hchan
        - hchan is protected by by mutex lock

  * Buffer Full
    - Lets assume with the above code itself,
      - with G1, the buffer is is full with 3,2,1
      - now G1 wants to send another value on to the buffer, A value 4 after aquiring a lock
      - since the buffer is full, G1  will blocked and will wait for receive.
        - how does block happen ?
          - G1 creates a a sudo G struct, and `G` element will hold the refference to the goroutine G1 and the value to be sent will be saved in the `elem` field. This structure is enqued inot the `sendq` list.
          - buffer-full-send.png
          - G1 calls to scheduler to a call to gopark(), the scheduler will move G1 out of execution on the OS thread, and other Go routines on the LRQ will get scheduled on the OS thread.
          - Now G2 comes along and tries to receive data from the channel.
            - It first aquires a lock,
            - deques the element(1) from the queue and copies the elemnt into the variable `v` and pops the waiting G1 from the `sendq` and enques the value saved in the `elem` field
            - important, G2 is the one which enques the value on to the buffer on which G1 was blocked. This is for optimisation as G1 does not have to do any channel operation again.
            - Once enques is done by G2, it will set the state of G1 as runnable again, this is done by G2 calling goready(G1) function
            - Then G1 is moved to the runnable state and then gets added to the local run queue. G1 will get scheduled by the OS thread when it gets a chance to execute.
            - buffer-full-recv.png


  * Buffer Empty
    - consider same prog above, assume G2 will execute first anf try to receive from a empty channel.
    - The buffer is empty, and G2 is called a receive on an empty channel and blocked
    - Now G2 will create a sudo struct, and enques itself onto to `recvq` with `G` pointing to G2 and `elem` pointing to the stack variable `v`.
    - buffer-empty-recv.png
    - G2 call to scheduler to a call to gopark(), the scheduler moves the G2 from the os thread and does a context switching with next go routines in the LRQ.
    - Now G1 comes along and tries to send a value on to the channel, first it checks if there are any go routines waiting in the `recvq` of the channel and it finds G2.
    - Now G1 copies the value into the variable of the G2 stack directly. This is the only scenario where one goroutine acces the stack of another goroutine. This is done for performance reasons so that G2 does not have to come later and perform any channel operations, and there is one fewer memory copy.(G1-accessing-g2-stack.png)
    - G1 pops G2 forom the `recvq` and puts it in the runnable state, by calling the goready(g2) function.
    - Now G2 moves back to local run queue, and gets schedules by the scheduler on the OS thread for a chance to run.

  * Unbuffered channel
    - Send on a unbuffered channel
      - When sender goroutine wants to send a values
      - if there is corresponding receiver waiting in recvq
        - Sender will write the value directly into the receiver goroutine stack variable
        - Sender goroutine puts the receiver goroutinr back to runnable state.
      - If there is no receiver goroutine in recvq
        - Sender gets parked into the sendq
        - Data is saved in the elem field in sudog struct.
        - Receiver comes and copies the data
        - Puts the sender to runnable state again
    - Receive on unbuffered channel
      - Receiver goroutine wants to receive value
      - if it find a goroutine in waiting sendq
        - Receiver copies the value in elem field to its variable
        - puts the sender goroutine to runnable state.
      - if there was no sender goroutine in sendq
        - Receiver gets parked into the `recvq`
        - Reference to variable is saved in the elem filed in sudog struct
        - Sender comes and copies the data directly to receiver stack variable
        - puts the receiver to runnable state.

  * Summary
    - hchan struct represents the channel
    - it contains circular ring buffer and mutex lock
    - Goroutines that gets blocked on send or recv are parked in sendq or recvq
    - Go scheduler moves the blocked goroutines, out of OS thread
    - Once channel operation is complete, goroutine is moved back to local run queue

* Select
  - select-scenario.png
  - useful when we can do operation on channel which ever is ready and don't worry about the order?
    -
      ```
        select {
          case <-ch1:
              // block of statements
          case <-ch2:
              // block of statements
          case ch3 <- struct{}{}
            // block of statements
        }
      ```
    - select statements is kike a switch
    - Each cases specifies communication
    - All channel operation are considered simultaneously, to see if any of them are ready
    - Select waits until some case is ready to proceed, or else the entire select statement is going to block until some case is ready for the communication.
    - when one of the channel is ready, that operation will proceed, and execute the associated block of statements.
    - if multiple channels are ready, select will pick one at random and execute
    - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/03-select/01-select/main.go
    - Select is also very helpful in implementing
      - timeouts
        - select waits untik there is event on channel or until timeout is reached.
        - The time.After function takes in a time duration and return a channel that will run in background as a go routine and will send the current time after the duration you provide it.
        ```
          select {
            case v := <-ch:
              fmt.Println(v)
            case <- time.After(3 * time.Second):
               fmt.Println("timeout")
          }
        ```
        - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/03-select/02-select/main.go
      - non-blocking communication(with default case)
        - send or receive on a channel, but avoid blocking if channel is not ready
        - Default allows you to exit a select block without blocking.
        ```
          select {
            case m := <- ch:
              fmt.println("received message", m)
            default:
              fmt.println(" no received message")
          }
        ```
        - https://github.com/raj23manj/go-concurrency-exercises-ind/blob/master/01-exercise-solution/03-select/03-select/main.go
      - Some scenarios to consider
        - Empty select statement will block forever
          `select{}`
        - Select on nil channel will block forever.
          ```
            var ch chan string
            // not make here hence nil
            select {
              case v := <-ch:
              case ch <- v:
            }
          ```




