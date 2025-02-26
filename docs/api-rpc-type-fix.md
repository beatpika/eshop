# API 和 RPC 层类型转换问题修复方案

## 问题描述

在 `app/api/biz/service/login.go` 文件中发现类型不匹配问题：

1. API 层使用 `github.com/beatpika/eshop/app/api/hertz_gen/basic/user.UserLoginReq` 类型调用 RPC 客户端的 Login 方法，但该方法期望接收 `github.com/beatpika/eshop/rpc_gen/kitex_gen/user.LoginReq` 类型。

2. RPC 客户端返回 `github.com/beatpika/eshop/rpc_gen/kitex_gen/user.LoginResp` 类型，但代码试图将其直接赋值给 API 层的 `github.com/beatpika/eshop/app/api/hertz_gen/basic/user.UserLoginResp` 类型。

这是因为 API 层（Hertz）和 RPC 层（Kitex）使用了不同的 protobuf 定义生成的结构体：

- API 层使用 `idl/api/handler_user.proto`
- RPC 层使用 `idl/user.proto`

## 解决方案

需要在 `app/api/biz/service/login.go` 中进行以下修改：

1. 添加必要的导入：
```go
kitex_user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
"github.com/beatpika/eshop/app/api/hertz_gen/common"
```

2. 修改 Run 方法中的 RPC 调用代码，进行类型转换：
```go
// 将 API 请求转换为 RPC 请求
rpcReq := &kitex_user.LoginReq{
    Email:    req.Email,
    Password: req.Password,
}

// 调用 RPC
rpcResp, err := rpc.UserClient.Login(h.Context, rpcReq)
if err != nil {
    return nil, err
}

// 将 RPC 响应转换为 API 响应
resp = &user.UserLoginResp{
    Base: &common.BaseResp{
        Code:    0,
        Message: "success",
    },
    UserId: rpcResp.UserId,
    // 如果需要 token，可以在这里设置
}
```

## 实施步骤

1. 切换到 Code 模式
2. 应用上述代码修改
3. 测试登录功能确保修复生效

## 后续建议

1. 考虑在 API 和 RPC 层之间添加统一的类型转换工具函数
2. 在后续开发中尽量统一 API 和 RPC 层的数据结构定义，减少不必要的类型转换