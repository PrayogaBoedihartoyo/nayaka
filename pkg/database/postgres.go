package database

import (
    "database/sql"

    _ "github.com/lib/pq"
)

func NewPostgresConnection(databaseURL string) (*sql.DB, error) {
    return sql.Open("postgres", databaseURL)
}