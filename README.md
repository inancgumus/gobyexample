# Go by Example: Programmer's Guide to Idiomatic and Testable Code

**Unlock Go's unique perspective on program design and start writing simple, maintainable, efficient, and testable code through hands-on, realistic project examples.**

You can read the book at https://mng.bz/DwMy.

<a href="https://mng.bz/DwMy"><image src="https://github.com/inancgumus/gobyexample/assets/621232/ff3efc27-86bc-427b-bd98-db56f9be09e7" width="40%" alt="Read Go by Example" /></a>

## What is this book about?
* Example projects based on real-world scenarios.
* Writing idiomatic, efficient, and testable Go code from scratch.
* Writing code that is easy to read, maintain, and extend.
* Demonstrating Go's unique features and using them as they were intended.
* Go patterns, best practices, and common pitfalls to avoid.

## Who is this book for?
* Expert programmers proficient in another Object-Oriented language and want to learn idiomatic Go.
* Gophers who want to improve their skills and knowledge for writing idiomatic, efficient, and testable code.

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

#### Topics covered
- Exploring idiomatic guidelines with packages
- Satisfying the standard library interfaces
- Writing and running tests using the tools provided by Go
- Using table-driven testing and subtests to improve maintainability
- Writing example tests to generate runnable documentation


## [Chapter 3: Test coverage and optimization](all-listings/03-test-coverage-and-optimization/README.md)



This chapter covers the importance of test coverage and performance optimization.

#### Topics covered
- Measuring test coverage to see what percentage of code is tested
- Optimizing for performance using benchmarks and profiling
- Parallel testing to reduce test run time and detect data race issues


## [Chapter 4: Command-line interfaces](all-listings/04-command-line-interfaces/README.md)



This chapter uses the standard library alone to write idiomatic command-line programs.

#### Topics covered
- Structuring and writing user-friendly command-line tools
- Parsing command-line arguments and flags
- Exploring the standard library's `os` and `flag` packages
- Extending the `flag` package with custom types


## [Chapter 5: Dependency injection](all-listings/05-dependency-injection/README.md)



This chapter explores how to decouple code from its dependencies to improve testability and maintainability.

#### Topics covered
- Exploring challenges to testability
- Improving testability by isolating dependencies
- Using dependency injection techniques for testability


## [Chapter 6: Synchronous APIs for concurrency](all-listings/06-synchronous-apis-for-concurrency/README.md)



This chapter focuses on designing and structuring synchronous APIs that treat concurrency as an implementation detail.

#### Topics covered
- Designing synchronous APIs that treat concurrency as an implementation detail
- Standardizing the processing of sequences of values using push iterators
- Structuring concurrent processing with the pipeline pattern


## [Chapter 7: Responsive and efficient programs](all-listings/07-responsive-and-efficient-programs/README.md)



This chapter focuses on writing responsive and efficient programs, using HTTP as a case study.

#### Topics covered
- Propagating cancellation signals using the `context` package
- Efficiently processing HTTP using the `net/http` package
- Efficiently consuming byte streams using the `io` package
- Interface embedding and composition mechanics
- Idiomatic testing techniques using the `RoundTripper` interface
- Testing against a test HTTP server using the `httptest` package


## [Chapter 8: Structuring packages and services](all-listings/08-structuring-packages-and-services/README.md)



This chapter focuses on package organization and structuring that avoids import cycles, using HTTP services as a case study.

#### Topics covered
- Effectively organizing and structuring packages to avoid import cycles
- Simplifying the use of types by making their zero value useful
- Protecting against unsafe concurrent use of shared state with mutexes
- Implementing and testing HTTP servers with the `http` and `httptest` packages
- Recording HTTP responses with custom test helpers to streamline testing


## [Chapter 9: Composition patterns](all-listings/09-composition-patterns/README.md)



This chapter focuses on composition patterns to demonstrate how to design packages that offer reusable and composable functionality.

#### Topics covered
- Effectively using composition patterns and functional programming techniques
- Leveraging field embedding and method forwarding to reuse functionality
- Understanding optional interfaces, method hiding issues, and type assertions
- Propagating request-scoped values across package APIs Using `Context`
- Implementing custom `slog.Handler`s to log extra attributes automatically
- Wrapping interfaces to modify, extend, and intercept existing behavior
- Extracting implicit behavior using anonymous interfaces


## [Chapter 10: Polymorphic storage](all-listings/10-polymorphic-storage/README.md)



This chapter explains using implicit interfaces to decouple storage from its consumers and highlights consumer-defined interfaces.

#### Topics covered
- Interacting with SQL databases
- Understanding the driver pattern
- Integration-testing external services
- Leveraging consumer-driven interface approach


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

#### Topics covered
- Goroutines and their usage
- Using `WaitGroup` to synchronize goroutines
- Unbuffered channels for communication between goroutines
- Closing channels to signal completion
- Using `select` to handle multiple channel operations
- Buffered channels for decoupling goroutines
- Many other widely used concurrency patterns


## [Appendix F: Self-referential options](all-listings/af-self-referential-options/README.md)



This appendix covers Rob Pike's self-referential options pattern and compares it with other approaches.


## [Appendix G: Cross-compiling Go programs](all-listings/ag-cross-compiling-go-programs/README.md)



This appendix shows how to cross-compile Go programs for different operating systems and architectures.
<!-- LISTING LINKS END -->
