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
* Coordinanting tasks -> WaitGroups
* Shared memory -> Mutexes
