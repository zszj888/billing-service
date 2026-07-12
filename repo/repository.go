package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samz/billing/cache"
	"github.com/samz/billing/domain"
)

type BillRepository struct {
	db    *sql.DB
	cache *cache.BillCache
}

func (repo BillRepository) Save(c context.Context, bill *domain.BillEntity) error {
	query := `
INSERT INTO bill (
    bill_no,
    title,
    description,
    amount,
    currency,
    status,
    due_date,
    paid_at,
    created_at,
    updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := repo.db.ExecContext(
		c,
		query,
		bill.BillNo,
		bill.Title,
		bill.Description,
		bill.Amount,
		bill.Currency,
		bill.Status,
		bill.DueDate,
		bill.PaidAt,
	)
	return err
}

func NewBillRepository(db *sql.DB, rdb *redis.Client) *BillRepository {
	return &BillRepository{
		db:    db,
		cache: cache.NewBillCache(rdb, 5*time.Minute),
	}
}

func (repo BillRepository) GetOneBill(ctx context.Context, id int64) (bill domain.BillEntity, _ error) {
	bill, err := repo.cache.GetBill(ctx, id)
	if err == nil {
		return bill, nil
	}
	if !errors.Is(err, redis.Nil) {
		log.Printf("redis GET error for bill %d: %v", id, err)
	}

	err = repo.db.QueryRowContext(ctx,
		"SELECT id, bill_no, title, description, amount, currency, status, due_date, paid_at, created_at, updated_at "+
			"FROM bill WHERE id = ?", id).Scan(
		&bill.ID,
		&bill.BillNo,
		&bill.Title,
		&bill.Description,
		&bill.Amount,
		&bill.Currency,
		&bill.Status,
		&bill.DueDate,
		&bill.PaidAt,
		&bill.CreatedAt,
		&bill.UpdatedAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return domain.BillEntity{}, domain.ErrNotFound
	case err != nil:
		return domain.BillEntity{}, err
	default:
		if cacheErr := repo.cache.SetBill(ctx, bill); cacheErr != nil {
			log.Printf("redis SET error for bill %d: %v", id, cacheErr)
		}
		return bill, nil
	}
}
