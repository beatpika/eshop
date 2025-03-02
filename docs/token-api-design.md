# Token API 设计文档

## 概述

Token服务提供token的生成、验证、刷新和撤销功能。本文档描述了对外提供的RESTful API接口设计。

## Proto文件设计

### 基础结构
```protobuf
syntax = "proto3";

package frontend.auth;

import "api/api.proto";
import "api/common.proto";

option go_package = "basic/token";
```

### 消息定义

#### 1. 生成Token
```protobuf
message GenerateTokenReq {
    int32 user_id = 1 [(api.form) = "user_id"];
}

message GenerateTokenResp {
    common.BaseResp base = 1;
    string access_token = 2;
    string refresh_token = 3;
    int64 expires_in = 4;  // token过期时间（秒）
}
```

#### 2. 刷新Token
```protobuf
message RefreshTokenReq {
    string refresh_token = 1 [(api.form) = "refresh_token"];
}

message RefreshTokenResp {
    common.BaseResp base = 1;
    string access_token = 2;
    string refresh_token = 3;
    int64 expires_in = 4;
}
```

#### 3. 验证Token
```protobuf
message VerifyTokenReq {
    string token = 1 [(api.form) = "token"];
}

message VerifyTokenResp {
    common.BaseResp base = 1;
    bool is_valid = 2;
    int32 user_id = 3;  // 如果token有效，返回对应的用户ID
}
```

#### 4. 撤销Token
```protobuf
message RevokeTokenReq {
    string token = 1 [(api.form) = "token"];
}

message RevokeTokenResp {
    common.BaseResp base = 1;
}
```

### 服务定义
```protobuf
service TokenHandler {
    rpc GenerateToken(GenerateTokenReq) returns (GenerateTokenResp) {
        option (api.post) = "/token/generate";
    }
    
    rpc RefreshToken(RefreshTokenReq) returns (RefreshTokenResp) {
        option (api.post) = "/token/refresh";
    }
    
    rpc VerifyToken(VerifyTokenReq) returns (VerifyTokenResp) {
        option (api.post) = "/token/verify";
    }
    
    rpc RevokeToken(RevokeTokenReq) returns (RevokeTokenResp) {
        option (api.post) = "/token/revoke";
    }
}
```

## API 接口说明

### 1. 生成Token
- 路径: `/token/generate`
- 方法: POST
- 描述: 根据用户ID生成新的access token和refresh token
- 权限: 需要验证请求来源是否合法

### 2. 刷新Token
- 路径: `/token/refresh`
- 方法: POST
- 描述: 使用refresh token获取新的access token
- 权限: 需要验证refresh token的有效性

### 3. 验证Token
- 路径: `/token/verify`
- 方法: POST
- 描述: 验证token的有效性，并返回相关的用户信息
- 权限: 公开接口

### 4. 撤销Token
- 路径: `/token/revoke`
- 方法: POST
- 描述: 使指定的token失效
- 权限: 需要验证请求来源是否合法