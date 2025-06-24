# Chapter 8: Structuring packages and services

This chapter focuses on package organization and structuring that avoids import cycles, using HTTP services as a case study.

#### Topics covered
- Effectively organizing and structuring packages to avoid import cycles
- Simplifying the use of types by making their zero value useful
- Protecting against unsafe concurrent use of shared state with mutexes
- Implementing and testing HTTP servers with the `http` and `httptest` packages
- Recording HTTP responses with custom test helpers to streamline testing

## Sections

### Section 1: Organizing and structuring packages
- This section has no listings.
### Section 2: Core
- [Listing 8.1: Standardized errors](../all-listings/08-structuring-packages-and-services/01-standardized-errors.md)
- [Listing 8.2: Core types](../all-listings/08-structuring-packages-and-services/02-core-types.md)
- [Listing 8.3: Validation](../all-listings/08-structuring-packages-and-services/03-validation.md)
- [Listing 8.4: `Shorten`](../all-listings/08-structuring-packages-and-services/04-shorten.md)
- [Listing 8.5: `Shortener` service](../all-listings/08-structuring-packages-and-services/05-shortener-service.md)
- [Listing 8.6: `Shortener.Resolve`](../all-listings/08-structuring-packages-and-services/06-shortenerresolve.md)
- [Listing 8.7: `RWMutex` protection](../all-listings/08-structuring-packages-and-services/07-rwmutex-protection.md)
### Section 3: HTTP
- [Listing 8.8: A health check handler](../all-listings/08-structuring-packages-and-services/08-a-health-check-handler.md)
- [Listing 8.9: Running an HTTP server](../all-listings/08-structuring-packages-and-services/09-running-an-http-server.md)
### Section 4: HTTP Handlers
- [Listing 8.10: `Shorten` handler](../all-listings/08-structuring-packages-and-services/10-shorten-handler.md)
- [Listing 8.11: `Resolve` handler](../all-listings/08-structuring-packages-and-services/11-resolve-handler.md)
- [Listing 8.12: Errors to HTTP status codes](../all-listings/08-structuring-packages-and-services/12-errors-to-http-status-codes.md)
### Section 5: Routing
- [Listing 8.13: Routing with `ServeMux`](../all-listings/08-structuring-packages-and-services/13-routing-with-servemux.md)
### Section 6: Timeouts
- [Listing 8.14: Setting server timeouts](../all-listings/08-structuring-packages-and-services/14-setting-server-timeouts.md)
### Section 7: Testing
- [Listing 8.15: Testing a handler](../all-listings/08-structuring-packages-and-services/15-testing-a-handler.md)
- [Listing 8.16: Adding a test helper](../all-listings/08-structuring-packages-and-services/16-adding-a-test-helper.md)
- [Listing 8.17: Testing with a test helper](../all-listings/08-structuring-packages-and-services/17-testing-with-a-test-helper.md)
### Section 8: Simplicity
- This section has no listings.
### Section 9: Extra (not included in the book)
- [Listing 8.18: Add `linkc` instructions (not included in the book)](https://github.com/inancgumus/gobyexample/blob/73acb3755b4ff3af94e6ce90b55de4792acb0f1a/link/cmd/linkc/README.md)
