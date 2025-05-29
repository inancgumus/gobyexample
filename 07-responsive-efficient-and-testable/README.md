# Chapter 7: Responsive, efficient, and testable

This chapter focuses on writing responsive, efficient, and testable code, using HTTP as a case study.

#### Topics covered
- Propagating cancellation signals using the `context` package.
- Efficiently processing HTTP using the `net/http` package.
- Efficiently consuming byte streams using the `io` package.
- Interface embedding and composition mechanics of Go.
- Idiomatic testing techniques using the `RoundTripper` interface.
- Testing against a fake HTTP server using the `httptest` package.

## Sections

### Section 1: Cancellation propagation
- [Listing 7.1: Cancelable pipeline](../all-listings/07-responsive-efficient-and-testable/01-cancelable-pipeline.md)
- [Listing 7.2: Deriving Contexts](../all-listings/07-responsive-efficient-and-testable/02-deriving-contexts.md)
- [Listing 7.3: Catching interruption signals](../all-listings/07-responsive-efficient-and-testable/03-catching-interruption-signals.md)
### Section 2: HTTP and efficient I/O operations
- [Listing 7.4: Processing a request and response](../all-listings/07-responsive-efficient-and-testable/04-processing-a-request-and-response.md)
### Section 3: Optimization
- [Listing 7.5: Optimizing http.Client](../all-listings/07-responsive-efficient-and-testable/05-optimizing-httpclient.md)
### Section 4: Testing
- [Listing 7.6: Satisfying RoundTripper](../all-listings/07-responsive-efficient-and-testable/06-satisfying-roundtripper.md)
- [Listing 7.7: Testing with a fake RoundTripper](../all-listings/07-responsive-efficient-and-testable/07-testing-with-a-fake-roundtripper.md)
### Section 5: End-to-end testing
- [Listing 7.8: Testing SendN](../all-listings/07-responsive-efficient-and-testable/08-testing-sendn.md)
### Section 6: Exercises
- [Exercise 1: Solution](../all-listings/07-responsive-efficient-and-testable/09-exercise-1-solution.md)
