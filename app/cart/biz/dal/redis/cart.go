package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/beatpika/eshop/app/cart/conf"
	"github.com/redis/go-redis/v9"
)

const (
	cartPrefix = "cart:"
)

type CartItem struct {
	ProductID uint32 `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

// GetCartKey returns the Redis key for a user's cart
func GetCartKey(userID uint32) string {
	return fmt.Sprintf("%s%d", cartPrefix, userID)
}

// AddItem adds or updates an item in user's cart
func AddItem(ctx context.Context, userID uint32, item *CartItem) error {
	rdb := conf.GetRedisClient()
	key := GetCartKey(userID)

	// Get current cart
	items := make(map[uint32]*CartItem)
	data, err := rdb.Get(ctx, key).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}

	if err != redis.Nil {
		if err := json.Unmarshal(data, &items); err != nil {
			return err
		}
	}

	// Update or add item
	items[item.ProductID] = item

	// Save back to Redis
	data, err = json.Marshal(items)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, data, 0).Err()
}

// GetCart retrieves all items in user's cart
func GetCart(ctx context.Context, userID uint32) (map[uint32]*CartItem, error) {
	rdb := conf.GetRedisClient()
	key := GetCartKey(userID)

	items := make(map[uint32]*CartItem)
	data, err := rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return items, nil
	}
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// EmptyCart removes all items from user's cart
func EmptyCart(ctx context.Context, userID uint32) error {
	rdb := conf.GetRedisClient()
	key := GetCartKey(userID)

	return rdb.Del(ctx, key).Err()
}
