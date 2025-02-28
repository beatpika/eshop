package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// AccessTokenExpiry 访问令牌过期时间 - 2小时
	AccessTokenExpiry = 2 * time.Hour

	// RefreshTokenExpiry 刷新令牌过期时间 - 30天
	RefreshTokenExpiry = 30 * 24 * time.Hour
)

type JWTUtil struct {
	secretKey string
}

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID int32  `json:"user_id"`
	Role   int32  `json:"role"`
	Nonce  string `json:"nonce,omitempty"` // 添加nonce字段确保token唯一性
}

func NewJWTUtil(secretKey string) *JWTUtil {
	return &JWTUtil{
		secretKey: secretKey,
	}
}

// GenerateToken 生成JWT token
func (j *JWTUtil) GenerateToken(userID int32, role int32, expiry time.Duration) (string, error) {
	now := time.Now()
	nonce := fmt.Sprintf("%d", now.UnixNano()) // 使用纳秒级时间戳作为nonce

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		UserID: userID,
		Role:   role,
		Nonce:  nonce,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

// VerifyToken 验证并解析JWT token
func (j *JWTUtil) VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

// IsTokenExpired 检查token是否过期
func (j *JWTUtil) IsTokenExpired(claims *CustomClaims) bool {
	if claims.ExpiresAt == nil {
		return true
	}
	return claims.ExpiresAt.Before(time.Now())
}
