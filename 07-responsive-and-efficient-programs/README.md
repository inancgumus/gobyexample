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
- [Listing 7.1: Cancelable pipeline](../all-listings/07-responsive-and-efficient-programs/01-cancelable-pipeline.md)
- [Listing 7.2: Deriving `Context`s](../all-listings/07-responsive-and-efficient-programs/02-deriving-contexts.md)
- [Listing 7.3: Catching interruption signals](../all-listings/07-responsive-and-efficient-programs/03-catching-interruption-signals.md)
### Section 3: HTTP and efficient I/O
- [Listing 7.4: Processing a request and response](../all-listings/07-responsive-and-efficient-programs/04-processing-a-request-and-response.md)
### Section 4: Optimization
- [Listing 7.5: Optimizing `http.Client`](../all-listings/07-responsive-and-efficient-programs/05-optimizing-httpclient.md)
### Section 5: Testing
- [Listing 7.6: Satisfying `RoundTripper`](../all-listings/07-responsive-and-efficient-programs/06-satisfying-roundtripper.md)
- [Listing 7.7: Testing with a fake `RoundTripper`](../all-listings/07-responsive-and-efficient-programs/07-testing-with-a-fake-roundtripper.md)
### Section 6: HTTP testing
- [Listing 7.8: Testing `SendN`](../all-listings/07-responsive-and-efficient-programs/08-testing-sendn.md)
### Section 7: Exercises
- [Exercise 1: Solution](../all-listings/07-responsive-and-efficient-programs/09-exercise-1-solution.md)
