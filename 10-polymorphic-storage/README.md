# Chapter 10: Polymorphic storage

This chapter explains using implicit interfaces to decouple storage from its consumers and highlights consumer-defined interfaces.

#### Topics covered
- Interacting with SQL databases
- Understanding the driver pattern
- Integration-testing external services
- Leveraging consumer-driven interface approach

## Sections

### Section 1: Interacting with SQL databases
- [Listing 10.1: Registering the SQLite driver](../all-listings/10-polymorphic-storage/01-registering-the-sqlite-driver.md)
- [Listing 10.2: Dialing the database](../all-listings/10-polymorphic-storage/02-dialing-the-database.md)
- [Listing 10.3: Defining the SQL schema](../all-listings/10-polymorphic-storage/03-defining-the-sql-schema.md)
- [Listing 10.4: Applying the schema](../all-listings/10-polymorphic-storage/04-applying-the-schema.md)
### Section 2: Database-backed service
- [Listing 10.5: Shortening links](../all-listings/10-polymorphic-storage/05-shortening-links.md)
- [Listing 10.6: Adding a test database helper](../all-listings/10-polymorphic-storage/06-adding-a-test-database-helper.md)
- [Listing 10.7: Testing `Shortener.Shorten`](../all-listings/10-polymorphic-storage/07-testing-shortenershorten.md)
- [Listing 10.8: Detecting constraint errors](../all-listings/10-polymorphic-storage/08-detecting-constraint-errors.md)
- [Listing 10.9: Detecting duplicate keys](../all-listings/10-polymorphic-storage/09-detecting-duplicate-keys.md)
- [Listing 10.10: Resolving links](../all-listings/10-polymorphic-storage/10-resolving-links.md)
### Section 3: Valuer and Scanner
- [Listing 10.11: `Valuer` and `Scanner`](../all-listings/10-polymorphic-storage/11-valuer-and-scanner.md)
- [Listing 10.12: Using `base64String`](../all-listings/10-polymorphic-storage/12-using-base64string.md)
### Section 4: Implicit interfaces
- [Listing 10.13: Declaring interfaces](../all-listings/10-polymorphic-storage/13-declaring-interfaces.md)
- [Listing 10.14: Integrating `sqlite.Shortener`](../all-listings/10-polymorphic-storage/14-integrating-sqliteshortener.md)
