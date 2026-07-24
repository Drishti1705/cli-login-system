package models

import "time"

type User struct {
	ID               int
	Username         string
	Password         string
	TOTPSecret       string
	TwoFactorEnabled bool
	FailedAttempts   int
	LockedUntil      *time.Time
	CreatedAt        time.Time
	LastLogin        *time.Time
}