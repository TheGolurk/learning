# Concurrency

## concurrency and parallelism

You can't have parallelism within concurrency.


For example making 3 tasks
Concurrency: Make 3 task but not simultaneously

Parallelism: Make 3 task at the same time


|Thread|Goroutine|
|:-----|:--------|
|Have own execution stack|Have own execution stack|
|Fixed stack space (around 1 MB)|Variable stack space (starts at @2 KB)|
|Managed by OS|Managed by Go runtime|

challenges with concurrency
* Coordinanting tasks -> WaitGroups -> chanels
* Shared memory -> Mutexes -> chanels

## sync.WaitGroup
A WaitGroup waits for a collection of goroutines to finish

## Mutex
A **mut**ual **ex**clusion lock

## see race running go 
```
$ go run --race main.go
```

## Mutexes
* Mutex: Ensures only one task can access
* RWMutex: Multiple Readers access protected code but only one task can write 


## Chanels
> "Don't communicate by sharing memory, share memory by communicating" **-Rob Pike**

Creating Chanels
```go
// create a channel
ch := make(chan int)

// create a buffered channel
ch := make(chan int, 5)


```


## Channel Types
* Bidirectional
* Send-only
* Receive-only

```go
ch := make(chan int) // created channels are always bidirectional
func myFunc(ch chan int) {...} // Bidirectional channel

func myFunc(ch chan<- int) {...} // send-only channel

func myFunc(ch <-chan int) {...} // receive-only channel
```


## Closing channels
* Closed via the built-in close function (close)
* Cannot check for closed channel!!
* Sending new message triggers a panic
* Receiving messages okay
	* if buffered, all buffered messages available
	* if unbuffered, or buffer empty, receive zero-value
* Use comma okay syntax to check


## Control flow
* If statements
* For loops
* Select statements


## Select Statement
```go
ch1 := make(chan int)
ch2 := make(chan string)

select {
	case i := <-ch1:
		// ....
	case ch2 <- "hello":
		// ...
	default:
		// use default case for non-blocking select
}
```
