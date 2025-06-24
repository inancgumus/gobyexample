# Chapter 3: Test coverage and optimization

This chapter covers the importance of test coverage and performance optimization.

#### Topics covered
- Measuring test coverage to see what percentage of code is tested
- Optimizing for performance using benchmarks and profiling
- Parallel testing to reduce test run time and detect data race issues

## Sections

### Section 1: Test coverage
- [Listing 3.1: Testing for edge cases](01-testing-for-edge-cases.md)
- [Listing 3.2: Adding an empty scheme test case](02-adding-an-empty-scheme-test-case.md)
- [Listing 3.3: Fix for empty schemes](03-fix-for-empty-schemes.md)
- [Listing 3.4: Testing `String` with a `nil` pointer](04-testing-string-with-a-nil-pointer.md)
- [Listing 3.5: Fix for a `nil` pointer](05-fix-for-a-nil-pointer.md)
- [Listing 3.6: Testing `String` with an empty URL](06-testing-string-with-an-empty-url.md)
- [Listing 3.7: Reassembling a URL](07-reassembling-a-url.md)
### Section 2: Benchmarking and optimization
- [Listing 3.8: Writing a benchmark for `String`](08-writing-a-benchmark-for-string.md)
- [Listing 3.9: Writing sub-benchmarks](09-writing-sub-benchmarks.md)
- [Listing 3.10: Optimizing the `String` method](10-optimizing-the-string-method.md)
### Section 3: Compiler optimizations
- This section has no listings.
### Section 4: Parallel testing
- [Listing 3.11: Parallel tests](11-parallel-tests.md)
- [Listing 3.12: Parallel subtests](12-parallel-subtests.md)
- [Listing 3.13: Data race](13-data-race.md)
