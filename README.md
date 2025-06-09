# Go by Example: Programmer's Guide to Idiomatic and Testable Code

You can read this book at https://www.manning.com/books/go-by-example.

<a href="https://www.manning.com/books/go-by-example"><image src="https://github.com/inancgumus/gobyexample/assets/621232/ff3efc27-86bc-427b-bd98-db56f9be09e7" width="40%" alt="Read Go by Example" /></a>

## What is this book about?
* Example projects based on real-world scenarios.
* Writing idiomatic and testable Go code from scratch.
* Writing code that is easy to read, maintain, and extend.
* Demonstrating Go's unique features and how to use them effectively.
* Go patterns, best practices, and common pitfalls.

## Who is this book for?
* Gophers who want to improve their skills and knowledge for writing idiomatic and testable code.
* Expert programmers proficient in another Object-Oriented language and want to learn Go.

> This book includes a **Go crash course** in Appendices.

## Table of Contents

> [!TIP]
> Click on the chapter headings to see the code listings.

<!-- LISTING LINKS START -->


## [Chapter 1: Getting started](all-listings/01-getting-started/README.md)



This chapter explores the book's goals, introduces Go, and showcases Go's prominent features without fully delving into Go's mechanics and idioms. It gets you started with Go and prepares you for the more advanced topics covered in the book.

#### Topics covered
- Why should you read this book?
- Importance of writing idiomatic and testable code
- Presenting Go's prominent features


## [Chapter 2: Idioms and testing](all-listings/02-idioms-and-testing/README.md)



This chapter focuses on writing idiomatic code and testing.

It provides practical examples to illustrate these concepts through the lens of the standard library.

#### Topics covered
- Exploring idiomatic guidelines with packages.
- Satisfying the standard library interfaces.
- Writing and running tests using the tools provided by Go.
- Using table-driven testing and subtests to improve maintainability.
- Writing example tests to generate runnable documentation.


## [Chapter 3: Test coverage and optimization](all-listings/03-test-coverage-and-optimization/README.md)



This chapter covers the importance of test coverage and performance optimization. It explores measuring test coverage, identifying untested code, and optimizing performance using built-in profiling tools.

#### Topics covered
- Measuring test coverage to see what percentage of code is tested.
- Optimizing for performance using benchmarks and profiling.
- Parallel testing to reduce test run time and detect data race issues.


## [Chapter 4: Command-line interfaces](all-listings/04-command-line-interfaces/README.md)



This chapter uses the standard library alone to write idiomatic command-line programs.

#### Topics covered
- Structuring and writing user-friendly command-line tools.
- Parsing command-line arguments and flags using the `os` and `flag` packages.
- Extending the `flag` package with custom value parsers.
- Validating flags after parsing to ensure correctness.
- Using positional arguments for mandatory arguments.


## [Chapter 5: Dependency injection](all-listings/05-dependency-injection/README.md)



This chapter explores how to decouple code from its dependencies to improve testability and maintainability. It illustrates the challenges of writing testable code and provides techniques for isolating dependencies using dependency injection.

#### Topics covered
- Exploring the challenges to testability.
- Improving testability by isolating dependencies.
- Dependency injection techniques for testability.


## [Chapter 6: Synchronous APIs for concurrency](all-listings/06-synchronous-apis-for-concurrency/README.md)



This chapter focuses on designing and structuring synchronous APIs that treat concurrency as an implementation detail.

#### Topics covered
- Designing synchronous APIs that treat concurrency as an implementation detail.
- Standardizing the processing of sequences of values using push iterators.
- Structuring concurrent processing with the pipeline pattern.


## [Chapter 7: Responsive and efficient programs](all-listings/07-responsive-and-efficient-programs/README.md)



This chapter focuses on writing responsive and efficient programs, using HTTP as a case study.

#### Topics covered
- Propagating cancellation signals using the `context` package.
- Efficiently processing HTTP using the `net/http` package.
- Efficiently consuming byte streams using the `io` package.
- Interface embedding and composition mechanics of Go.
- Idiomatic testing techniques using the `RoundTripper` interface.
- Testing against a fake HTTP server using the `httptest` package.


## [Chapter 8: Structuring packages and services](all-listings/08-structuring-packages-and-services/README.md)



This chapter focuses on package organization and structuring that avoids import cycles, using HTTP services as a case study.

#### Topics covered
- Effectively organizing and structuring packages to avoid import cycles.
- Simplifying the usage of types by making their zero value useful.
- Protecting against unsafe concurrent use of shared state with mutexes.
- Implementing and testing HTTP servers with the `http` and `httptest` packages.
- Recording HTTP responses with custom test helpers to streamline testing.


## [Chapter 9: Composition patterns](all-listings/09-composition-patterns/README.md)



This chapter focuses on composition patterns to demonstrate how to design packages that offer reusable and composable functionality.

#### Topics covered
- Composition patterns and functional programming techniques.
- Field embedding and method forwarding to provide extra functionality.
- Optional interfaces, method hiding issues, and type assertions.
- Using `Context` to propagate values across packages APIs.
- Implementing a custom `slog.Handler` to log extra log attributes.
- Interface wrapping techniques to change existing behavior.
- Effectively using anonymous interfaces for extracting behavior.


## [Chapter 10: Polymorphic storage](all-listings/10-polymorphic-storage/README.md)



This chapter explains using implicit interfaces to decouple storage from its consumers, interacting with SQL databases via Go's sql package, and effectively testing database-backed code. It highlights consumer-defined interfaces to improve modularity and composability.

#### Topics covered
- Using the `sql` package to interact with SQL databases.
- Idiomatic design patterns used in the `sql` and `sql/driver` packages.
- Integration testing database-backed code effectively.
- Decoupling storage through consumer-defined interfaces.
- Enhancing modularity and composability with focused interfaces.


## [Appendix A: Modules and packages](all-listings/aa-modules-and-packages/README.md)



This appendix introduces the concept of modules and packages.


## [Appendix B: Variables and pointers](all-listings/ab-variables-and-pointers/README.md)



This appendix covers the basics of variables and pointers.


## [Appendix C: Arrays, Slices, and Maps](all-listings/ac-arrays-slices-and-maps/README.md)



This appendix covers the basics of arrays, slices, and maps.


## [Appendix D: Object-oriented programming](all-listings/ad-object-oriented-programming/README.md)



This appendix covers the basics of object-oriented programming in Go.


## [Appendix E: Concurrent programming](all-listings/ae-concurrent-programming/README.md)



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


## [Appendix F: Self-referential options](all-listings/af-self-referential-options/README.md)



This appendix covers Rob Pike's self-referential options pattern and compares it with other options patterns.


## [Appendix G: Cross-compiling Go programs](all-listings/ag-cross-compiling-go-programs/README.md)



This appendix shows how to cross-compile Go programs for different operating systems and architectures.
<!-- LISTING LINKS END -->
