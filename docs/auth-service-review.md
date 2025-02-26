# Auth 服务设计问题和建议

## 现有问题

1. 命名和语义问题
   - `DeliverTokenByRPC` 和 `VerifyTokenByRPC` 名称中包含 "ByRPC" 是多余的，因为 gRPC 服务本身就是 RPC 调用
   - `DeliverToken` 不如使用 `GenerateToken` 或 `CreateToken` 更准确
   - 响应消息命名不一致（`DeliveryResp` vs `VerifyResp`）

2. 字段设计问题
   - `DeliverTokenReq` 中的 `user_id` 默认值设置为 0 可能导致误判
   - `VerifyTokenReq` 中的 token 默认值 "emtp" 不合适
   - `DeliveryResp` 中的 token 默认值 "emtp" 不合适
   - `VerifyResp` 使用 `res` 作为字段名不够清晰

3. 功能完整性问题
   - 缺少 token 过期时间设置
   - 缺少 token 刷新机制
   - 缺少用户角色/权限信息
   - 缺少错误处理机制
   - 验证响应过于简单，缺少具体的验证信息

## 改进建议

```protobuf
syntax="proto3";

package auth;

option go_package="/auth";

// 错误码枚举
enum ErrorCode {
    ERROR_CODE_UNSPECIFIED = 0;
    ERROR_CODE_INVALID_TOKEN = 1;
    ERROR_CODE_TOKEN_EXPIRED = 2;
    ERROR_CODE_USER_NOT_FOUND = 3;
    ERROR_CODE_PERMISSION_DENIED = 4;
}

// 用户角色枚举
enum UserRole {
    USER_ROLE_UNSPECIFIED = 0;
    USER_ROLE_USER = 1;
    USER_ROLE_MERCHANT = 2;
    USER_ROLE_ADMIN = 3;
}

service AuthService {
    // 生成新的访问令牌
    rpc GenerateToken(GenerateTokenRequest) returns (GenerateTokenResponse) {}
    
    // 验证令牌
    rpc VerifyToken(VerifyTokenRequest) returns (VerifyTokenResponse) {}
    
    // 刷新令牌
    rpc RefreshToken(RefreshTokenRequest) returns (RefreshTokenResponse) {}
    
    // 撤销令牌
    rpc RevokeToken(RevokeTokenRequest) returns (RevokeTokenResponse) {}
}

message GenerateTokenRequest {
    int32 user_id = 1;
    UserRole role = 2;
    repeated string permissions = 3; // 可选的额外权限列表
}

message GenerateTokenResponse {
    string access_token = 1;
    string refresh_token = 2;
    int64 expires_at = 3;     // Unix 时间戳，token过期时间
    ErrorCode error_code = 4;
    string error_message = 5;
}

message VerifyTokenRequest {
    string token = 1;
}

message VerifyTokenResponse {
    bool is_valid = 1;
    int32 user_id = 2;
    UserRole role = 3;
    repeated string permissions = 4;
    int64 expires_at = 5;      // Unix 时间戳，token过期时间
    ErrorCode error_code = 6;
    string error_message = 7;
}

message RefreshTokenRequest {
    string refresh_token = 1;
}

message RefreshTokenResponse {
    string access_token = 1;
    string refresh_token = 2;  // 可选的新刷新令牌
    int64 expires_at = 3;
    ErrorCode error_code = 4;
    string error_message = 5;
}

message RevokeTokenRequest {
    string token = 1;
}

message RevokeTokenResponse {
    bool success = 1;
    ErrorCode error_code = 2;
    string error_message = 3;
}
```

## 改进说明

1. 接口完整性
   - 添加了令牌刷新机制
   - 添加了令牌撤销功能
   - 增加了过期时间管理
   - 添加了用户角色和权限支持

2. 错误处理
   - 添加了统一的错误码枚举
   - 在响应中增加了错误信息字段
   - 提供了更详细的错误描述

3. 安全性增强
   - 支持访问令牌和刷新令牌分离
   - 加入了用户角色管理
   - 支持细粒度的权限控制

4. 字段改进
   - 移除了不恰当的默认值
   - 使用更清晰的字段命名
   - 添加了必要的元数据字段

5. 可扩展性
   - 通过枚举预留了错误码空间
   - 通过字符串数组支持灵活的权限定义
   - 考虑了多角色场景

## 实现建议

1. 令牌管理
   - 使用 JWT 作为令牌格式
   - 实现令牌黑名单机制
   - 使用 Redis 等缓存系统存储令牌状态

2. 安全建议
   - 令牌传输使用 HTTPS
   - 实现令牌轮换机制
   - 设置合理的令牌过期时间
   - 记录令牌操作审计日志

3. 性能优化
   - 使用缓存减少验证开销
   - 实现令牌预验证机制
   - 支持批量令牌操作