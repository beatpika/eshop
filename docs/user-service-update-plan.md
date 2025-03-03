# 用户服务更新实施计划

## 1. 更新 Proto 文件

### 1.1 修改 user.proto
```protobuf
syntax = "proto3";

package user;

option go_package = "/user";

service UserService {
    // 现有接口
    rpc Register(RegisterReq) returns (RegisterResp)
    rpc Login(LoginReq) returns (LoginResp)
    
    // 新增接口
    rpc GetUserInfo(GetUserInfoReq) returns (GetUserInfoResp)
    rpc UpdateUserInfo(UpdateUserInfoReq) returns (UpdateUserInfoResp)
    rpc UpdatePassword(UpdatePasswordReq) returns (UpdatePasswordResp)
    rpc UpdatePhone(UpdatePhoneReq) returns (UpdatePhoneResp)
    rpc DeactivateAccount(DeactivateAccountReq) returns (DeactivateAccountResp)
}

message RegisterReq {
    string email = 1;
    string username = 2;
    string password = 3;
    string phone = 4;
}

message RegisterResp {
    int32 user_id = 1;
    string username = 2;
}

message LoginReq {
    string email = 1;
    string password = 2;
}

message LoginResp {
    int32 user_id = 1;
    string username = 2;
    string email = 3;
    string avatar = 4;
}

message GetUserInfoReq {
    int32 user_id = 1;
}

message GetUserInfoResp {
    UserInfo user = 1;
}

message UpdateUserInfoReq {
    int32 user_id = 1;
    string username = 2;
    string avatar = 3;
    string address = 4;
}

message UpdateUserInfoResp {
    bool success = 1;
}

message UpdatePasswordReq {
    int32 user_id = 1;
    string old_password = 2;
    string new_password = 3;
}

message UpdatePasswordResp {
    bool success = 1;
}

message UpdatePhoneReq {
    int32 user_id = 1;
    string phone = 2;
    string verify_code = 3;
}

message UpdatePhoneResp {
    bool success = 1;
}

message DeactivateAccountReq {
    int32 user_id = 1;
    string password = 2;
}

message DeactivateAccountResp {
    bool success = 1;
}

message UserInfo {
    int32 user_id = 1;
    string username = 2;
    string email = 3;
    string phone = 4;
    string avatar = 5;
    string address = 6;
    int32 status = 7;
    int64 created_at = 8;
}
```

### 1.2 修改 handler_user.proto
```protobuf
syntax = "proto3";

package frontend.auth;

import "api/api.proto";
import "api/common.proto";

option go_package = "basic/user";

service UserHandler {
    rpc login(UserLoginReq) returns (UserLoginResp) {
        option (api.post) = "/user/login";
    }
    rpc register(UserRegisterReq) returns (UserRegisterResp) {
        option (api.post) = "/user/register";
    }
    rpc logout(UserLogoutReq) returns (UserLogoutResp) {
        option (api.post) = "/user/logout";
    }
    rpc getUserInfo(GetUserInfoReq) returns (GetUserInfoResp) {
        option (api.get) = "/user/info";
    }
    rpc updateUserInfo(UpdateUserInfoReq) returns (UpdateUserInfoResp) {
        option (api.put) = "/user/info";
    }
    rpc updatePassword(UpdatePasswordReq) returns (UpdatePasswordResp) {
        option (api.put) = "/user/password";
    }
    rpc updatePhone(UpdatePhoneReq) returns (UpdatePhoneResp) {
        option (api.put) = "/user/phone";
    }
    rpc deactivateAccount(DeactivateAccountReq) returns (DeactivateAccountResp) {
        option (api.post) = "/user/deactivate";
    }
}
```

## 2. 更新数据库模型

### 2.1 修改 user.go
```go
type User struct {
    gorm.Model
    Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
    Username       string `gorm:"type:varchar(255) not null"`
    PasswordHashed string `gorm:"type:varchar(255) not null"`
    Phone         string `gorm:"type:varchar(20)"`
    Avatar        string `gorm:"type:varchar(255)"`
    Address       string `gorm:"type:text"`
    Status        int    `gorm:"type:tinyint;default:1"` // 1:正常 2:禁用 3:待验证
}
```

## 3. 代码生成步骤

1. 生成 RPC 代码：
```bash
cd rpc_gen
go generate ./...
```

2. 生成 API 代码：
```bash
cd app/api
hz update -idl ../../idl/api/handler_user.proto
```

## 4. 实现新接口

1. 在 user 服务中实现新的 RPC 方法
2. 在 API 网关中实现新的 HTTP handler
3. 添加必要的中间件（参数验证、认证等）
4. 实现错误处理

## 5. 测试计划

1. 单元测试：
   - 新增的 service 方法
   - 密码强度验证
   - 手机号验证
   
2. 集成测试：
   - API 接口测试
   - 错误处理测试
   - 并发测试

## 6. 部署计划

1. 数据库迁移：
   - 创建新的数据库迁移文件
   - 测试环境验证
   - 生产环境迁移

2. 服务部署：
   - 测试环境部署
   - 验证新功能
   - 生产环境部署

## 7. 监控和告警

1. 添加新的监控指标：
   - 用户注册成功率
   - 登录成功率
   - API 响应时间
   - 错误率统计

2. 设置告警规则：
   - 高错误率告警
   - 响应时间过长告警
   - 并发量超限告警