package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"testing"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

// Dial connects to SQLite and applies the schema for convenience.
func Dial(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging: %w", err)
	}
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, fmt.Errorf("applying schema: %w", err)
	}
	return db, nil
}

// DialTestDB connects to a unique in-memory SQLite database.
func DialTestDB(tb testing.TB) *sql.DB {
	tb.Helper()

	dsn := fmt.Sprintf(
		"file:%s?mode=memory&cache=shared",
		tb.Name(),
	)
	db, err := Dial(tb.Context(), dsn)
	if err != nil {
		tb.Fatalf("DialTestDB: %v", err)
	}
	tb.Cleanup(func() {
		if err := db.Close(); err != nil {
			tb.Logf("DialTestDB: closing: %v", err)
		}
	})

	return db
}
