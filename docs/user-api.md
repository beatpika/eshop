# User API 文档

本文档描述了用户服务提供的所有 RESTful API 接口。

## 通用说明

### Base URL
```
http://localhost:8080
```

### 请求格式
所有 POST/PUT 请求使用 `application/x-www-form-urlencoded` 格式。

### 认证方式
除注册和登录外，所有接口都需要在请求头中携带 token：
```
Authorization: Bearer your-token-here
```

### 响应格式
```json
{
    "status_code": 200,      // 状态码，200 表示成功
    "status_message": "Success" // 状态描述
    // 其他字段...
}
```

## API 列表

### 1. 用户注册
- **路径**: `/user/register`
- **方法**: POST
- **参数**:
  - email: 邮箱
  - username: 用户名
  - password: 密码
  - password_confirm: 确认密码
  - phone: 手机号

**请求示例**:
```bash
curl -X POST 'http://localhost:8080/user/register' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'email=test@example.com' \
-d 'username=testuser' \
-d 'password=password@123' \
-d 'password_confirm=password@123' \
-d 'phone=13800138000'
```

### 2. 用户登录
- **路径**: `/user/login`
- **方法**: POST
- **参数**:
  - email: 邮箱
  - password: 密码

**请求示例**:
```bash
curl -X POST 'http://localhost:8080/user/login' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'email=test@example.com' \
-d 'password=password@123'
```

### 3. 用户登出
- **路径**: `/user/logout`
- **方法**: POST
- **认证**: 需要

**请求示例**:
```bash
curl -X POST 'http://localhost:8080/user/logout' \
-H 'Authorization: Bearer your-token-here' \
-H 'Content-Type: application/x-www-form-urlencoded'
```

### 4. 获取用户信息
- **路径**: `/user/info`
- **方法**: GET
- **认证**: 需要

**请求示例**:
```bash
curl -X GET 'http://localhost:8080/user/info' \
-H 'Authorization: Bearer your-token-here'
```

### 5. 更新用户信息
- **路径**: `/user/info`
- **方法**: PUT
- **认证**: 需要
- **参数**:
  - username: 新用户名
  - avatar: 头像URL
  - address: 地址

**请求示例**:
```bash
curl -X PUT 'http://localhost:8080/user/info' \
-H 'Authorization: Bearer your-token-here' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'username=newname' \
-d 'avatar=https://example.com/avatar.jpg' \
-d 'address=北京市朝阳区'
```

### 6. 更新密码
- **路径**: `/user/password`
- **方法**: PUT
- **认证**: 需要
- **参数**:
  - old_password: 旧密码
  - new_password: 新密码

**请求示例**:
```bash
curl -X PUT 'http://localhost:8080/user/password' \
-H 'Authorization: Bearer your-token-here' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'old_password=password123' \
-d 'new_password=newpassword123'
```

### 7. 更新手机号
- **路径**: `/user/phone`
- **方法**: PUT
- **认证**: 需要
- **参数**:
  - phone: 新手机号
  - verify_code: 验证码

**请求示例**:
```bash
curl -X PUT 'http://localhost:8080/user/phone' \
-H 'Authorization: Bearer your-token-here' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'phone=13900139000' \
-d 'verify_code=123456'
```

### 8. 注销账号
- **路径**: `/user/deactivate`
- **方法**: POST
- **认证**: 需要
- **参数**:
  - password: 密码（确认）

**请求示例**:
```bash
curl -X POST 'http://localhost:8080/user/deactivate' \
-H 'Authorization: Bearer your-token-here' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'password=password123'
```

## 状态码说明

| 状态码 | 描述 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未认证或认证失败 |
| 403 | 无权限访问 |
| 500 | 服务器内部错误 |