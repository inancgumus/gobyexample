# Chapter 10: Polymorphic storage

This chapter explains using implicit interfaces to decouple storage from its consumers and highlights consumer-defined interfaces.

#### Topics covered
- Interacting with SQL databases
- Understanding the driver pattern
- Integration-testing external services
- Leveraging consumer-driven interface approach

## Sections

### Section 1: Interacting with SQL databases
- [Listing 10.1: Registering the SQLite driver](01-registering-the-sqlite-driver.md)
- [Listing 10.2: Dialing the database](02-dialing-the-database.md)
- [Listing 10.3: Defining the SQL schema](03-defining-the-sql-schema.md)
- [Listing 10.4: Applying the schema](04-applying-the-schema.md)
### Section 2: Database-backed service
- [Listing 10.5: Shortening links](05-shortening-links.md)
- [Listing 10.6: Adding a test database helper](06-adding-a-test-database-helper.md)
- [Listing 10.7: Testing `Shortener.Shorten`](07-testing-shortenershorten.md)
- [Listing 10.8: Detecting constraint errors](08-detecting-constraint-errors.md)
- [Listing 10.9: Detecting duplicate keys](09-detecting-duplicate-keys.md)
- [Listing 10.10: Resolving links](10-resolving-links.md)
### Section 3: Valuer and Scanner
- [Listing 10.11: `Valuer` and `Scanner`](11-valuer-and-scanner.md)
- [Listing 10.12: Using `base64String`](12-using-base64string.md)
### Section 4: Implicit interfaces
- [Listing 10.13: Declaring interfaces](13-declaring-interfaces.md)
- [Listing 10.14: Integrating `sqlite.Shortener`](14-integrating-sqliteshortener.md)
