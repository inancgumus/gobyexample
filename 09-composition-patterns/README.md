# Chapter 9: Composition patterns

This chapter focuses on composition patterns to demonstrate how to design packages that offer reusable and composable functionality.

#### Topics covered
- Effectively using composition patterns and functional programming techniques
- Leveraging field embedding and method forwarding to reuse functionality
- Understanding optional interfaces, method hiding issues, and type assertions
- Propagating request-scoped values across package APIs Using `Context`
- Implementing custom `slog.Handler`s to log extra attributes automatically
- Wrapping interfaces to modify, extend, and intercept existing behavior
- Extracting implicit behavior using anonymous interfaces

## Sections

### Section 1: Middleware pattern
- [Listing 9.1: Logger middleware](../all-listings/09-composition-patterns/01-logger-middleware.md)
- [Listing 9.2: Activating the middleware](../all-listings/09-composition-patterns/02-activating-the-middleware.md)
### Section 2: Logging responses
- [Listing 9.3: `Duration` middleware](../all-listings/09-composition-patterns/03-duration-middleware.md)
- [Listing 9.4: Response recording](../all-listings/09-composition-patterns/04-response-recording.md)
- [Listing 9.5: Integrating response recording](../all-listings/09-composition-patterns/05-integrating-response-recording.md)
### Section 3: Interceptor pattern
- [Listing 9.6: Intercepting method calls](../all-listings/09-composition-patterns/06-intercepting-method-calls.md)
- [Listing 9.7: `StatusCode` middleware](../all-listings/09-composition-patterns/07-statuscode-middleware.md)
- [Listing 9.8: Integrating `StatusCode`](../all-listings/09-composition-patterns/08-integrating-statuscode.md)
### Section 4: Optional interface pattern
- [Listing 9.9: Unwrap](../all-listings/09-composition-patterns/09-unwrap.md)
### Section 5: `Context` value propagation pattern
- [Listing 9.10: Package `traceid`](../all-listings/09-composition-patterns/10-package-traceid.md)
- [Listing 9.11: `traceid` middleware](../all-listings/09-composition-patterns/11-traceid-middleware.md)
- [Listing 9.12: `LogHandler`](../all-listings/09-composition-patterns/12-loghandler.md)
- [Listing 9.13: `WithAttrs` and `WithGroup`](../all-listings/09-composition-patterns/13-withattrs-and-withgroup.md)
- [Listing 9.14: Integrating trace IDs](../all-listings/09-composition-patterns/14-integrating-trace-ids.md)
### Section 6: Handler chaining pattern
- [Listing 9.15: A chainable custom handler type](../all-listings/09-composition-patterns/15-a-chainable-custom-handler-type.md)
- [Listing 9.16: `Responder` type](../all-listings/09-composition-patterns/16-responder-type.md)
- [Listing 9.17: `Error` response helper](../all-listings/09-composition-patterns/17-error-response-helper.md)
- [Listing 9.18: `Redirect` response helper](../all-listings/09-composition-patterns/18-redirect-response-helper.md)
- [Listing 9.19: `Text` response helper](../all-listings/09-composition-patterns/19-text-response-helper.md)
- [Listing 9.20: `Responder` helper](../all-listings/09-composition-patterns/20-responder-helper.md)
- [Listing 9.21: Chainable handlers](../all-listings/09-composition-patterns/21-chainable-handlers.md)
### Section 7: Encoding and decoding
- [Listing 9.22: JSON response helper](../all-listings/09-composition-patterns/22-json-response-helper.md)
- [Listing 9.23: Decoding JSON](../all-listings/09-composition-patterns/23-decoding-json.md)
- [Listing 9.24: Speaking JSON](../all-listings/09-composition-patterns/24-speaking-json.md)
### Section 8: Wrapping and unwrapping
- [Listing 9.25: `MaxBytesReader` helper](../all-listings/09-composition-patterns/25-maxbytesreader-helper.md)
- [Listing 9.26: Integrating `MaxBytesReader`](../all-listings/09-composition-patterns/26-integrating-maxbytesreader.md)
- [Listing 9.27: Unwrapping `ResponseWriter`s](../all-listings/09-composition-patterns/27-unwrapping-responsewriters.md)
### Section 9: Outro
- This section has no listings.
### Section 10: Extra (not included in the book)
- [Listing 9.28: Extra (not included in the book)](../all-listings/09-composition-patterns/28-extra.md)
