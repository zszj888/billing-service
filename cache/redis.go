package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samz/billing/domain"
)

type BillCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewBillCache(client *redis.Client, ttl time.Duration) *BillCache {
	return &BillCache{client: client, ttl: ttl}
}

func (c *BillCache) GetBill(ctx context.Context, id int64) (domain.BillEntity, error) {
	key := fmt.Sprintf("bill:%d", id)
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.BillEntity{}, err
	}
	var bill domain.BillEntity
	if err := json.Unmarshal(data, &bill); err != nil {
		return domain.BillEntity{}, err
	}
	return bill, nil
}

func (c *BillCache) SetBill(ctx context.Context, bill domain.BillEntity) error {
	key := fmt.Sprintf("bill:%d", bill.ID)
	data, err := json.Marshal(bill)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, c.ttl).Err()
}

func (c *BillCache) InvalidateBill(ctx context.Context, id int64) error {
	key := fmt.Sprintf("bill:%d", id)
	return c.client.Del(ctx, key).Err()
}
