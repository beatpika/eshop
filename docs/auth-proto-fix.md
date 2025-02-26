# auth.proto 中的 "emtp" 问题修复

## 问题描述

在当前的 auth.proto 文件中，发现以下问题：

```protobuf
message VerifyTokenReq {
    string token = "emtp";    // 错误：不应该设置默认值
}

message DeliveryResp {
    string token = "emtp";    // 错误：不应该设置默认值
}
```

## 存在的问题

1. 在 Protocol Buffers 中，字符串类型的默认值应该是空字符串 ""
2. 设置 token 的默认值是一个安全隐患
3. "emtp" 没有明确的业务含义，可能是错误的占位符

## 建议修复

应该移除这些默认值，修改为：

```protobuf
message VerifyTokenReq {
    string token = 1;    // 移除默认值，使用字段编号
}

message DeliveryResp {
    string token = 1;    // 移除默认值，使用字段编号
}
```

## 安全建议

1. 不要在 proto 文件中设置敏感字段的默认值
2. token 应该在运行时动态生成
3. 确保 token 的生成使用安全的随机数生成器
4. 实现 token 的有效期检查
5. 考虑实现 token 轮换机制