# Token RPC 服务实现计划

## 1. 系统架构

### 1.1 整体架构
- 基于 Kitex RPC 框架实现 Token 服务
- 使用 Redis 存储 token 相关信息
- 采用分层架构: handler -> service -> dal

### 1.2 核心组件
1. RPC 接口层 (handler.go)
   - 已完成基础框架
   - 实现了接口路由到对应 service

2. 业务逻辑层 (biz/service/*)
   - 已搭建四个核心服务框架
   - 需要实现具体业务逻辑

3. 数据访问层 (biz/dal/*)
   - 需要实现 Redis 操作接口
   - 需要设计 token 存储结构

## 2. 核心功能实现计划

### 2.1 Token 生成服务 (GenerateToken)
1. 实现内容:
   - JWT token 生成器
   - Access token 和 Refresh token 的生成
   - Token 信息存储到 Redis

2. 关键设计:
   ```
   // Redis Key 结构
   access_token:{token} -> {user_id, role, expires_at}
   refresh_token:{token} -> {user_id, access_token}
   user_tokens:{user_id} -> Set<access_token>
   ```

### 2.2 Token 验证服务 (VerifyToken)
1. 实现内容:
   - JWT token 验证
   - Redis token 信息验证
   - 用户权限验证

2. 验证流程:
   - 验证 JWT 签名
   - 检查 Redis 中 token 是否存在
   - 验证过期时间
   - 返回 token 包含的用户信息

### 2.3 Token 刷新服务 (RefreshToken)
1. 实现内容:
   - 验证 refresh token 有效性
   - 生成新的 access token
   - 可选生成新的 refresh token

2. 刷新流程:
   - 验证 refresh token 存在性和有效性
   - 删除旧的 access token
   - 生成新的 access token
   - 更新 Redis 存储

### 2.4 Token 撤销服务 (RevokeToken)
1. 实现内容:
   - 支持撤销单个 token
   - 支持撤销用户所有 token

2. 撤销流程:
   - 从 Redis 删除 token 信息
   - 从用户 token 集合中移除
   - 处理关联的 refresh token

## 3. 技术方案

### 3.1 Token 设计
1. Access Token:
   - 使用 JWT 格式
   - 包含: user_id, role, exp
   - 有效期: 2小时

2. Refresh Token:
   - 使用 UUID 格式
   - 有效期: 30天
   - 支持自动续期

### 3.2 存储设计
1. Redis 数据结构:
   - String: 存储 token 详细信息
   - Set: 存储用户的 token 集合
   - Hash: 存储 refresh token 映射

2. 过期策略:
   - 使用 Redis TTL 机制
   - Token 过期自动清理
   - 定期清理过期数据

## 4. 安全考虑

1. Token 安全:
   - 使用强密钥加密
   - 避免敏感信息存储
   - 实现 token 黑名单机制

2. 防护措施:
   - Rate limiting
   - Token 重放攻击防护
   - 异常监控和告警

## 5. 实现步骤

1. 第一阶段:
   - 实现基础 JWT 工具类
   - 实现 Redis 操作接口
   - 完成 Generate Token 服务

2. 第二阶段:
   - 实现 Verify Token 服务
   - 实现 token 缓存机制
   - 添加基础监控

3. 第三阶段:
   - 实现 Refresh Token 服务
   - 实现 Revoke Token 服务
   - 完善错误处理

4. 第四阶段:
   - 添加完整的测试用例
   - 实现性能优化
   - 完善监控和告警

## 6. 测试计划

1. 单元测试:
   - 各个 service 的功能测试
   - Redis 操作测试
   - JWT 工具类测试

2. 集成测试:
   - RPC 接口测试
   - 多服务互操作测试
   - 性能和负载测试

3. 安全测试:
   - Token 伪造测试
   - 并发安全测试
   - 过期处理测试