# Product API Documentation

## 商品管理接口

### 创建商品

创建一个新的商品，包括基本信息和SKU列表。

```bash
curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试商品",
    "description": "这是一个测试商品",
    "category_id": 1,
    "images": [
      "http://example.com/image1.jpg",
      "http://example.com/image2.jpg"
    ],
    "price": 9900,
    "skus": [
      {
        "specs": {
          "color": "红色",
          "size": "XL"
        },
        "price": 9900,
        "stock": 100,
        "code": "TEST001"
      }
    ]
  }'
```

### 更新商品

更新商品的基本信息。

```bash
curl -X PUT http://localhost:8080/product/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "更新后的商品名称",
    "description": "更新后的商品描述",
    "category_id": 2,
    "images": [
      "http://example.com/new-image1.jpg"
    ],
    "price": 8800
  }'
```

### 获取商品详情

获取单个商品的详细信息。

```bash
curl -X GET http://localhost:8080/product/1
```

### 获取商品列表

获取商品列表，支持分页。

```bash
curl -X GET 'http://localhost:8080/products?page_size=10&page=1'
```

### 删除商品

删除指定商品。

```bash
curl -X DELETE http://localhost:8080/product/1
```

### 更新商品状态

更新商品的上架/下架状态。

```bash
curl -X PUT http://localhost:8080/product/1/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": 2
  }'
```

## SKU管理接口

### 创建SKU

为指定商品创建新的SKU。

```bash
curl -X POST http://localhost:8080/product/1/sku \
  -H "Content-Type: application/json" \
  -d '{
    "specs": {
      "color": "蓝色",
      "size": "M"
    },
    "price": 9900,
    "stock": 50,
    "code": "TEST002"
  }'
```

### 更新SKU

更新SKU的信息。

```bash
curl -X PUT http://localhost:8080/sku/1 \
  -H "Content-Type: application/json" \
  -d '{
    "specs": {
      "color": "绿色",
      "size": "L"
    },
    "price": 8800,
    "code": "TEST003"
  }'
```

### 删除SKU

删除指定SKU。

```bash
curl -X DELETE http://localhost:8080/sku/1
```

### 更新SKU库存

更新SKU的库存数量。

```bash
curl -X PUT http://localhost:8080/sku/1/stock \
  -H "Content-Type: application/json" \
  -d '{
    "stock": 200
  }'
```

## 分类管理接口

### 创建分类

创建新的商品分类。

```bash
curl -X POST http://localhost:8080/category \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试分类",
    "parent_id": 0,
    "level": 1,
    "sort_order": 1
  }'
```

### 更新分类

更新分类信息。

```bash
curl -X PUT http://localhost:8080/category/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "更新后的分类名称",
    "sort_order": 2
  }'
```

### 删除分类

删除指定分类。

```bash
curl -X DELETE http://localhost:8080/category/1
```

### 获取分类列表

获取指定父分类下的子分类列表。

```bash
curl -X GET 'http://localhost:8080/categories?parent_id=0'
```

### 获取分类树

获取完整的分类树结构。

```bash
curl -X GET http://localhost:8080/categories/tree
```

## 商品搜索接口

### 搜索商品

根据关键词搜索商品，支持分类过滤和排序。

```bash
curl -X GET 'http://localhost:8080/products/search?keyword=测试&category_id=1&page_size=10&page=1&sort_by=1'
```

### 获取分类下的商品

获取指定分类下的商品列表，支持分页和排序。

```bash
curl -X GET 'http://localhost:8080/categories/1/products?page_size=10&page=1&sort_by=1'
```

## 响应格式

所有接口返回的数据格式统一如下：

```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  },
  "data": {
    // 具体的响应数据
  }
}
```

### 状态码说明

- 0: 成功
- 400: 请求参数错误
- 401: 未授权
- 403: 禁止访问
- 404: 资源不存在
- 500: 服务器内部错误