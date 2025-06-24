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
- [Listing 9.1: Logger middleware](01-logger-middleware.md)
- [Listing 9.2: Activating the middleware](02-activating-the-middleware.md)
### Section 2: Logging responses
- [Listing 9.3: `Duration` middleware](03-duration-middleware.md)
- [Listing 9.4: Response recording](04-response-recording.md)
- [Listing 9.5: Integrating response recording](05-integrating-response-recording.md)
### Section 3: Interceptor pattern
- [Listing 9.6: Intercepting method calls](06-intercepting-method-calls.md)
- [Listing 9.7: `StatusCode` middleware](07-statuscode-middleware.md)
- [Listing 9.8: Integrating `StatusCode`](08-integrating-statuscode.md)
### Section 4: Optional interface pattern
- [Listing 9.9: Unwrap](09-unwrap.md)
### Section 5: `Context` value propagation pattern
- [Listing 9.10: Package `traceid`](10-package-traceid.md)
- [Listing 9.11: `traceid` middleware](11-traceid-middleware.md)
- [Listing 9.12: `LogHandler`](12-loghandler.md)
- [Listing 9.13: `WithAttrs` and `WithGroup`](13-withattrs-and-withgroup.md)
- [Listing 9.14: Integrating trace IDs](14-integrating-trace-ids.md)
### Section 6: Handler chaining pattern
- [Listing 9.15: A chainable custom handler type](15-a-chainable-custom-handler-type.md)
- [Listing 9.16: `Responder` type](16-responder-type.md)
- [Listing 9.17: `Error` response helper](17-error-response-helper.md)
- [Listing 9.18: `Redirect` response helper](18-redirect-response-helper.md)
- [Listing 9.19: `Text` response helper](19-text-response-helper.md)
- [Listing 9.20: `Responder` helper](20-responder-helper.md)
- [Listing 9.21: Chainable handlers](21-chainable-handlers.md)
### Section 7: Encoding and decoding
- [Listing 9.22: JSON response helper](22-json-response-helper.md)
- [Listing 9.23: Decoding JSON](23-decoding-json.md)
- [Listing 9.24: Speaking JSON](24-speaking-json.md)
### Section 8: Wrapping and unwrapping
- [Listing 9.25: `MaxBytesReader` helper](25-maxbytesreader-helper.md)
- [Listing 9.26: Integrating `MaxBytesReader`](26-integrating-maxbytesreader.md)
- [Listing 9.27: Unwrapping `ResponseWriter`s](27-unwrapping-responsewriters.md)
### Section 9: Outro
- This section has no listings.
### Section 10: Extra (not included in the book)
- [Listing 9.28: Extra (not included in the book)](28-extra.md)
