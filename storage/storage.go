package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Storage struct {
	db *sql.DB
}

// Init DB func for my app for processing mobile appeals
func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatal("Failed to create database", err)
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT NOT NULL,
		address TEXT NOT NULL,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL,
		permissions TEXT NOT NULL
	)`)
	if err != nil {
		return nil, fmt.Errorf("Failed to created table users: %w", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS appeals(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		theme TEXT NOT NULL,
		body TEXT NOT NULL,
		status TEXT NOT NULL,
		tariff_id INTEGER,
		FOREIGN KEY (tariff_id) REFERENCES tariffs (id),
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users (id),
		created_at TEXT
	)`)
	if err != nil {
		return nil, fmt.Errorf("Failed to created table appeals: %w", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS tariffs(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL
	)`)
	if err != nil {
		return nil, fmt.Errorf("Failed to created table tariffs: %w", err)
	}
	return &Storage{db: db}, nil
}

// Helper function for processing requests to DB
func (s *Storage) execQuery(ctx context.Context, query string, args ...any) error {
	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Printf("Error executing: %v", err)
		return err
	}
	return nil
}

// Add new record to table tariffs
func (s *Storage) AddTariff(title, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	query := "INSERT INTO tariffs(body, title) VALUES (?, ?)"
	defer cancel()

	return s.execQuery(ctx, query, title, body)
}

// Delete record from table tariffs
func (s *Storage) DelTariff(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	query := "DELETE FROM tariffs WHERE id = ?"
	defer cancel()

	return s.execQuery(ctx, query, id)
}

// Add new appeal
func (s *Storage) AddAppeal(theme, body string, tariffID, userID int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	query := "INSERT INTO appeals(theme, body, tariff_id, user_id) VALUES (?, ?, ?, ?)"
	defer cancel()

	return s.execQuery(ctx, query, theme, body, tariffID, userID)
}

// Delete appeal
func (s *Storage) DelAppeal(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	query := "DELETE FROM appeals WHERE id = ?"
	defer cancel()

	return s.execQuery(ctx, query, id)
}
