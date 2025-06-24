# Chapter 7: Responsive and efficient programs

This chapter focuses on writing responsive and efficient programs, using HTTP as a case study.

#### Topics covered
- Propagating cancellation signals using the `context` package
- Efficiently processing HTTP using the `net/http` package
- Efficiently consuming byte streams using the `io` package
- Interface embedding and composition mechanics
- Idiomatic testing techniques using the `RoundTripper` interface
- Testing against a test HTTP server using the `httptest` package

## Sections

### Section 1: Revisiting the concurrent pipeline
- This section has no listings.
### Section 2: Cancellation propagation
- [Listing 7.1: Cancelable pipeline](01-cancelable-pipeline.md)
- [Listing 7.2: Deriving `Context`s](02-deriving-contexts.md)
- [Listing 7.3: Catching interruption signals](03-catching-interruption-signals.md)
### Section 3: HTTP and efficient I/O
- [Listing 7.4: Processing a request and response](04-processing-a-request-and-response.md)
### Section 4: Optimization
- [Listing 7.5: Optimizing `http.Client`](05-optimizing-httpclient.md)
### Section 5: Testing
- [Listing 7.6: Satisfying `RoundTripper`](06-satisfying-roundtripper.md)
- [Listing 7.7: Testing with a fake `RoundTripper`](07-testing-with-a-fake-roundtripper.md)
### Section 6: HTTP testing
- [Listing 7.8: Testing `SendN`](08-testing-sendn.md)
### Section 7: Exercises
- [Exercise 1: Solution](09-exercise-1-solution.md)
