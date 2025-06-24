# Chapter 6: Synchronous APIs for concurrency

This chapter focuses on designing and structuring synchronous APIs that treat concurrency as an implementation detail.

#### Topics covered
- Designing synchronous APIs that treat concurrency as an implementation detail
- Standardizing the processing of sequences of values using push iterators
- Structuring concurrent processing with the pipeline pattern

## Sections

### Section 1: Overview
- This section has no listings.
### Section 2: Foundations
- [Listing 6.1: Implementing `Result`](01-implementing-result.md)
- [Listing 6.2: Simulating HTTP requests](02-simulating-http-requests.md)
### Section 3: Iterators
- [Listing 6.3: Implementing the `Results` iterator](03-implementing-the-results-iterator.md)
- [Listing 6.4: Returning a push iterator](04-returning-a-push-iterator.md)
- [Listing 6.5: `Summary` type](05-summary-type.md)
- [Listing 6.6: Summarizing the results](06-summarizing-the-results.md)
- [Listing 6.7: Testing `Summarize`](07-testing-summarize.md)
### Section 4: Options
- [Listing 6.8: Providing options](08-providing-options.md)
- [Listing 6.9: `SendN` with options](09-sendn-with-options.md)
### Section 5: Integration
- [Listing 6.10: Printing a summary](10-printing-a-summary.md)
- [Listing 6.11: Integrating the HIT client](11-integrating-the-hit-client.md)
### Section 6: Concurrent pipeline pattern
- This section has no listings.
### Section 7: Producer stage
- [Listing 6.12: Implementing the producer stage](12-implementing-the-producer-stage.md)
- [Listing 6.13: Integrating the producer stage](13-integrating-the-producer-stage.md)
### Section 8: Throttler stage
- [Listing 6.14: Implementing the throttler stage](14-implementing-the-throttler-stage.md)
- [Listing 6.15: Integrating the throttler stage](15-integrating-the-throttler-stage.md)
### Section 9: Dispatcher stage
- [Listing 6.16: Implementing the dispatcher](16-implementing-the-dispatcher.md)
- [Listing 6.17: Integrating the dispatcher stage](17-integrating-the-dispatcher-stage.md)
- [Listing 6.18: Updating the iterator to run the pipeline](18-updating-the-iterator-to-run-the-pipeline.md)
