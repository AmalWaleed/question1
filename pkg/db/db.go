// db.go

package main

import (
    "context"
    "database/sql"
    "time"

    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/jmoiron/sqlx"
)

// DB is a wrapper for database operations
type DB struct {
    db *sqlx.DB
}

// NewDB creates a new DB instance
func NewDB(pool *pgxpool.Pool) *DB {
    return &DB{db: sqlx.NewDb(pool, "pgx")}
}

// CreateUser creates a new user in the database
func (db *DB) CreateUser(ctx context.Context, user User) (int, error) {
    var userID int
    err := db.db.QueryRowContext(ctx, createUser, user.Name, user.PhoneNumber, user.OTP, user.OTPExpiration).Scan(&userID)
    return userID, err
}

// GenerateOTP generates OTP for a user in the database
func (db *DB) GenerateOTP(ctx context.Context, phoneNumber, otp string, expirationTime time.Time) (int, error) {
    var userID int
    err := db.db.QueryRowContext(ctx, generateOTP, otp, expirationTime, phoneNumber).Scan(&userID)
    return userID, err
}

// VerifyOTP verifies OTP for a user in the database
func (db *DB) VerifyOTP(ctx context.Context, phoneNumber, otp string) (bool, error) {
    var count int
    err := db.db.QueryRowContext(ctx, verifyOTP, phoneNumber, otp).Scan(&count)
    return count > 0, err
}
