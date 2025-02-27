package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// Token过期时间
	AccessTokenExpiry  = 2 * time.Hour
	RefreshTokenExpiry = 30 * 24 * time.Hour
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID int32 `json:"user_id"`
	Role   int32 `json:"role"`
}

type JWTUtil struct {
	secretKey []byte
}

// NewJWTUtil 从配置中获取密钥创建JWT工具
func NewJWTUtil(secretKey string) *JWTUtil {
	return &JWTUtil{
		secretKey: []byte(secretKey),
	}
}

// GenerateToken 生成JWT token
func (j *JWTUtil) GenerateToken(userID int32, role int32, expiry time.Duration) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
		UserID: userID,
		Role:   role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// VerifyToken 验证JWT token
func (j *JWTUtil) VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// IsTokenExpired 检查token是否过期
func (j *JWTUtil) IsTokenExpired(claims *TokenClaims) bool {
	if claims.ExpiresAt == nil {
		return true
	}
	return time.Now().After(claims.ExpiresAt.Time)
}

// GetExpirationTime 获取token过期时间
func (j *JWTUtil) GetExpirationTime(claims *TokenClaims) int64 {
	if claims.ExpiresAt == nil {
		return 0
	}
	return claims.ExpiresAt.Unix()
}
