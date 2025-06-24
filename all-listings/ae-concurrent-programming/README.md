# Appendix E: Concurrent programming

This appendix covers the basics of concurrent programming in Go.

#### Topics covered
- Goroutines and their usage
- Using `WaitGroup` to synchronize goroutines
- Unbuffered channels for communication between goroutines
- Closing channels to signal completion
- Using `select` to handle multiple channel operations
- Buffered channels for decoupling goroutines
- Many other widely used concurrency patterns

## Sections

### Section 1: Goroutines
- [Listing E.1: A sequential function](01-a-sequential-function.md)
- [Listing E.2: Running concurrently](02-running-concurrently.md)
### Section 2: Using WaitGroup
- [Listing E.3: `WaitGroup`](03-waitgroup.md)
- [Listing E.4: A safer `WaitGroup`](04-a-safer-waitgroup.md)
- [Listing E.5: Using `SafeGroup`](05-using-safegroup.md)
### Section 3: Unbuffered channels
- [Listing E.6: Sending and receiving](06-sending-and-receiving.md)
### Section 4: Closing channels
- [Listing E.7: Closing a channel](07-closing-a-channel.md)
- [Listing E.8: Ranging over a channel](08-ranging-over-a-channel.md)
- [Listing E.9: Close signal coordination](09-close-signal-coordination.md)
### Section 5: Using the select statement
- [Listing E.10: Selecting channels](10-selecting-channels.md)
- [Listing E.11: Default case](11-default-case.md)
### Section 6: Buffered channels
- [Listing E.12: Buffered channel](12-buffered-channel.md)
- [Listing E.13: Select with timeout](13-select-with-timeout.md)
- [Listing E.14: Counting semaphore](14-counting-semaphore.md)
