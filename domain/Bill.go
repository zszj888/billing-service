package domain

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("bill not found")

type BillEntity struct {
	ID          int64
	BillNo      string
	Title       string
	Description string
	Amount      float64
	Currency    string
	Status      BillStatus
	DueDate     time.Time
	PaidAt      *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BillStatus int8

const (
	BillPending BillStatus = iota
	BillPaid
	BillCancelled
)
