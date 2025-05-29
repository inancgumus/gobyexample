# Listing 10.3: Defining the SQL schema

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/ebd0224b014739ad8e9def951c946194be52ef3b/link) / [sqlite](https://github.com/inancgumus/gobyexample/blob/ebd0224b014739ad8e9def951c946194be52ef3b/link/sqlite) / [schema.sql](https://github.com/inancgumus/gobyexample/blob/ebd0224b014739ad8e9def951c946194be52ef3b/link/sqlite/schema.sql)

```go
CREATE TABLE IF NOT EXISTS links ( 
    short_key VARCHAR(16) PRIMARY KEY, 
    uri       TEXT NOT NULL 
);
```

