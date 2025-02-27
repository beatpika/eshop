package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// Redis key前缀
	accessTokenPrefix  = "access_token:"
	refreshTokenPrefix = "refresh_token:"
	userTokensPrefix   = "user_tokens:"
)

type TokenStore struct {
	client *redis.Client
}

func NewTokenStore(client *redis.Client) *TokenStore {
	return &TokenStore{client: client}
}

// SaveAccessToken 保存访问令牌
func (s *TokenStore) SaveAccessToken(ctx context.Context, token string, userID int32, role int32, expiresAt int64) error {
	key := accessTokenPrefix + token
	tokenData := map[string]interface{}{
		"user_id": userID,
		"role":    role,
		"expires": expiresAt,
	}

	// 保存token信息
	if err := s.client.HMSet(ctx, key, tokenData).Err(); err != nil {
		return fmt.Errorf("failed to save access token: %w", err)
	}

	// 设置过期时间
	expiration := time.Until(time.Unix(expiresAt, 0))
	if err := s.client.Expire(ctx, key, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set access token expiration: %w", err)
	}

	// 将token添加到用户的token集合
	userTokensKey := fmt.Sprintf("%s%d", userTokensPrefix, userID)
	if err := s.client.SAdd(ctx, userTokensKey, token).Err(); err != nil {
		return fmt.Errorf("failed to add token to user set: %w", err)
	}

	return nil
}

// SaveRefreshToken 保存刷新令牌
func (s *TokenStore) SaveRefreshToken(ctx context.Context, refreshToken, accessToken string, userID int32, expiresAt int64) error {
	key := refreshTokenPrefix + refreshToken
	tokenData := map[string]interface{}{
		"user_id":      userID,
		"access_token": accessToken,
	}

	// 保存refresh token信息
	if err := s.client.HMSet(ctx, key, tokenData).Err(); err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err)
	}

	// 设置过期时间
	expiration := time.Until(time.Unix(expiresAt, 0))
	if err := s.client.Expire(ctx, key, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set refresh token expiration: %w", err)
	}

	return nil
}

// GetAccessToken 获取访问令牌信息
func (s *TokenStore) GetAccessToken(ctx context.Context, token string) (map[string]string, error) {
	key := accessTokenPrefix + token
	data, err := s.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}
	if len(data) == 0 {
		return nil, redis.Nil
	}
	return data, nil
}

// GetRefreshToken 获取刷新令牌信息
func (s *TokenStore) GetRefreshToken(ctx context.Context, refreshToken string) (map[string]string, error) {
	key := refreshTokenPrefix + refreshToken
	data, err := s.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}
	if len(data) == 0 {
		return nil, redis.Nil
	}
	return data, nil
}

// DeleteAccessToken 删除访问令牌
func (s *TokenStore) DeleteAccessToken(ctx context.Context, token string, userID int32) error {
	key := accessTokenPrefix + token

	// 删除token
	if err := s.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete access token: %w", err)
	}

	// 从用户token集合中移除
	userTokensKey := fmt.Sprintf("%s%d", userTokensPrefix, userID)
	if err := s.client.SRem(ctx, userTokensKey, token).Err(); err != nil {
		return fmt.Errorf("failed to remove token from user set: %w", err)
	}

	return nil
}

// DeleteRefreshToken 删除刷新令牌
func (s *TokenStore) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	key := refreshTokenPrefix + refreshToken
	if err := s.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}
	return nil
}

// DeleteUserTokens 删除用户的所有令牌
func (s *TokenStore) DeleteUserTokens(ctx context.Context, userID int32) error {
	userTokensKey := fmt.Sprintf("%s%d", userTokensPrefix, userID)

	// 获取用户的所有token
	tokens, err := s.client.SMembers(ctx, userTokensKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get user tokens: %w", err)
	}

	// 删除所有access token
	for _, token := range tokens {
		key := accessTokenPrefix + token
		if err := s.client.Del(ctx, key).Err(); err != nil {
			return fmt.Errorf("failed to delete access token: %w", err)
		}
	}

	// 删除用户token集合
	if err := s.client.Del(ctx, userTokensKey).Err(); err != nil {
		return fmt.Errorf("failed to delete user tokens set: %w", err)
	}

	return nil
}
