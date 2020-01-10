# semaphore
--
    import "github.com/michilu/boilerplate/service/semaphore"


## Usage

#### func  Acquire

```go
func Acquire(ctx context.Context, n int) error
```
Acquire enters the semaphore a specified number of times, blocking only until
ctx is done.

#### func  GetCount

```go
func GetCount() int
```
GetCount returns current number of occupied entries in semaphore.

#### func  GetLimit

```go
func GetLimit() int
```
GetLimit returns current semaphore limit.

#### func  Init

```go
func Init(parallel int)
```

#### func  New

```go
func New(limit int) sem.Semaphore
```
New initializes a new instance of the Semaphore, specifying the maximum number
of concurrent entries.

#### func  Release

```go
func Release(n int) int
```
Release exits the semaphore a specified number of times and returns the previous
count.

#### func  SetParallel

```go
func SetParallel(i int) error
```
SetParallel sets a given number to the parallelism number.

#### func  TryAcquire

```go
func TryAcquire(n int) bool
```
TryAcquire acquires the semaphore without blocking.

#### type Semaphore

```go
type Semaphore = sem.Semaphore
```

Semaphore counting resizable semaphore synchronization primitive.
