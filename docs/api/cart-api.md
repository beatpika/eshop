# 购物车服务API文档

## 添加商品到购物车

### 请求
```http
POST /carts/:user_id/items
```

### 参数
| 参数名 | 类型 | 说明 | 位置 |
|--------|------|------|------|
| user_id | int64 | 用户ID | 路径参数 |
| product_id | int64 | 商品ID | 请求体 |
| quantity | int32 | 数量 | 请求体 |

### 响应
```json
{
  "base": {
    "status_code": 0,
    "status_message": "Success"
  }
}
```

### curl示例
```bash
curl -X POST 'http://localhost:8080/carts/1/items' \
-H 'Content-Type: application/json' \
-d '{
    "product_id": 1,
    "quantity": 2
}'
```

## 获取购物车信息

### 请求
```http
GET /carts/:user_id
```

### 参数
| 参数名 | 类型 | 说明 | 位置 |
|--------|------|------|------|
| user_id | int64 | 用户ID | 路径参数 |

### 响应
```json
{
  "base": {
    "status_code": 0,
    "status_message": "Success"
  },
  "cart": {
    "user_id": 1,
    "items": [
      {
        "product_id": 1,
        "product_name": "示例商品",
        "product_image": "http://example.com/image.jpg",
        "price": 1000,
        "quantity": 2,
        "subtotal": 2000
      }
    ],
    "total": 2000
  }
}
```

### curl示例
```bash
curl -X GET 'http://localhost:8080/carts/1'
```

## 清空购物车

### 请求
```http
POST /carts/:user_id/empty
```

### 参数
| 参数名 | 类型 | 说明 | 位置 |
|--------|------|------|------|
| user_id | int64 | 用户ID | 路径参数 |

### 响应
```json
{
  "base": {
    "status_code": 0,
    "status_message": "Success"
  }
}
```

### curl示例
```bash
curl -X POST 'http://localhost:8080/carts/1/empty'
```

## 错误码说明

| 状态码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 无效请求 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 未找到 |
| 500 | 内部错误 |

## 字段说明

### CartItem
| 字段名 | 类型 | 说明 |
|--------|------|------|
| product_id | int64 | 商品ID |
| product_name | string | 商品名称 |
| product_image | string | 商品图片URL |
| price | int64 | 单价(分) |
| quantity | int32 | 数量 |
| subtotal | int64 | 小计金额(分) |

### Cart
| 字段名 | 类型 | 说明 |
|--------|------|------|
| user_id | int64 | 用户ID |
| items | CartItem[] | 购物车商品列表 |
| total | int64 | 总金额(分) |