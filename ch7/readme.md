# GO Concurrency

- Go scheduler
- Channel : send data to goroutines, receive data from go routines or other special purpose.
- Shared memory


## Processes, threads and goroutines

- process: running binary file
- thread: as subset of a process
- goroutine: minimum Go entity

In practice, this means that a `process` can have `multiple threads` as well as lots of `goroutines`, 
whereas a goroutine needs the environment of a process to exist. So, to create a goroutine, 
you need to have a process with at least one thread. *The OS takes care of the process* and 
thread scheduling, while *Go creates the necessary threads* and the *developer creates 
the desired number of goroutines*.


## Go Scheduler
- Go runtime his its own scheduler.
- m:n scheduling: m goroutines are executed using n OS threads
- fork-join concurrency model
- work-stealing strategy 
- continuation-stealing


```go
Go Scheduler 是一個用於調度 Goroutines（Go 語言中輕量級的執行緒）的系統組件，它負責管理 Goroutines 的創建、調度和銷毀。Go Scheduler 基於協作式的調度器，這意味著它讓 Goroutines 在需要時自主地進行切換。

Go Scheduler 的基本機制如下：

Goroutines 被創建時，它們會被放入一個稱為 Goroutine 隊列的結構中。這個隊列存儲了所有需要運行的 Goroutines。

Go Scheduler 會選擇一個 Goroutine 開始執行，並將其放入一個稱為 P (Processor) 的結構中。P 負責執行 Goroutines。

在 P 中運行的 Goroutine 會在需要時進行切換。例如，如果 Goroutine 需要等待 I/O 完成，那麼 Go Scheduler 會選擇另一個 Goroutine 運行，直到 I/O 完成並且回到這個 Goroutine。

如果一個 Goroutine 需要長時間運行，Go Scheduler 會將其暫停，並將其放回 Goroutine 隊列中。然後，Go Scheduler 會選擇另一個 Goroutine 運行。

當一個 Goroutine 結束時，Go Scheduler 會將其從 P 中刪除，並從 Goroutine 隊列中移除。

簡單來說，Go Scheduler 會管理 Goroutines 的執行，確保它們在需要時運行，避免死鎖和競爭條件等問題。它是 Go 語言能夠高效地運行大量併發操作的關鍵。
```