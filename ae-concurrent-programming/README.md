# Appendix E: Concurrent programming

This appendix covers the basics of concurrent programming in Go.

It introduces goroutines, channels, `WaitGroup`, `select`, providing a foundation for writing concurrent programs.

Concurrency may involve parallelism, but they are not the same. Concurrency is about designing programs so that multiple tasks can make progress independently, even if they aren't running at the exact same time. It's more about managing how tasks interact and share resources than how they are physically executed. Parallelism, in contrast, is about executing multiple tasks truly simultaneously, typically across multiple CPUs. While this distinction may seem subtle, it's crucial for understanding how Go handles concurrency.

#### Topics covered
- Goroutines and their usage.
- Using `WaitGroup` to synchronize goroutines.
- Unbuffered channels for communication between goroutines.
- Closing channels to signal completion.
- Using `select` to handle multiple channel operations.
- Buffered channels for decoupling goroutines.
- Many other widely used concurrency patterns.

## Sections

### Section 1: Goroutines
- [Listing E.1: A sequential function](../all-listings/ae-concurrent-programming/01-a-sequential-function.md)
- [Listing E.2: Running concurrently](../all-listings/ae-concurrent-programming/02-running-concurrently.md)
### Section 2: WaitGroup
- [Listing E.3: WaitGroup](../all-listings/ae-concurrent-programming/03-waitgroup.md)
- [Listing E.4: A safer WaitGroup](../all-listings/ae-concurrent-programming/04-a-safer-waitgroup.md)
- [Listing E.5: Using SafeGroup](../all-listings/ae-concurrent-programming/05-using-safegroup.md)
### Section 3: Unbuffered channels
- [Listing E.6: Sending and receiving](../all-listings/ae-concurrent-programming/06-sending-and-receiving.md)
### Section 4: Closing channels
- [Listing E.7: Closing a channel](../all-listings/ae-concurrent-programming/07-closing-a-channel.md)
- [Listing E.8: Ranging over a channel](../all-listings/ae-concurrent-programming/08-ranging-over-a-channel.md)
- [Listing E.9: Close signal coordination](../all-listings/ae-concurrent-programming/09-close-signal-coordination.md)
### Section 5: Select
- [Listing E.10: Selecting channels](../all-listings/ae-concurrent-programming/10-selecting-channels.md)
- [Listing E.11: Default case](../all-listings/ae-concurrent-programming/11-default-case.md)
### Section 6: Buffered channels
- [Listing E.12: Buffered channel](../all-listings/ae-concurrent-programming/12-buffered-channel.md)
- [Listing E.13: Select with timeout](../all-listings/ae-concurrent-programming/13-select-with-timeout.md)
- [Listing E.14: Counting semaphore](../all-listings/ae-concurrent-programming/14-counting-semaphore.md)
