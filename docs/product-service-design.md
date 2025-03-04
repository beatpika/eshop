# 商品服务设计文档

## 1. 服务概述

商品服务（Product Service）是电商系统的核心服务之一，负责商品信息的管理、查询和库存控制。本文档详细说明商品服务的设计方案。

## 2. 核心功能

### 2.1 基础功能
- 商品的增删改查（CRUD）
- 商品分类管理
- SKU管理
- 库存管理
- 商品搜索和过滤

### 2.2 商品模型设计

#### 2.2.1 商品基础信息(Product)
```proto
message Product {
    int64 id = 1;                    // 商品ID
    string name = 2;                 // 商品名称
    string description = 3;          // 商品描述
    int64 category_id = 4;          // 分类ID
    repeated string images = 5;      // 商品图片URL列表
    int64 price = 6;                // 价格（单位：分）
    int32 status = 7;               // 商品状态：1-待上架 2-已上架 3-已下架
    repeated SKU skus = 8;          // SKU列表
    int64 created_at = 9;           // 创建时间
    int64 updated_at = 10;          // 更新时间
}
```

#### 2.2.2 SKU信息
```proto
message SKU {
    int64 id = 1;                    // SKU ID
    int64 product_id = 2;           // 商品ID
    map<string, string> specs = 3;   // 规格信息，如 {"color": "红色", "size": "XL"}
    int64 price = 4;                // SKU价格（单位：分）
    int32 stock = 5;                // 库存数量
    string code = 6;                // SKU编码
}
```

#### 2.2.3 商品分类
```proto
message Category {
    int64 id = 1;                    // 分类ID
    string name = 2;                 // 分类名称
    int64 parent_id = 3;            // 父分类ID，0表示一级分类
    int32 level = 4;                // 分类层级：1-一级分类 2-二级分类 3-三级分类
    int32 sort_order = 5;           // 排序序号
}
```

### 2.3 接口设计

#### 2.3.1 商品管理接口
- CreateProduct: 创建商品
- UpdateProduct: 更新商品信息
- GetProduct: 获取商品详情
- ListProducts: 获取商品列表
- DeleteProduct: 删除商品
- UpdateProductStatus: 更新商品状态（上架/下架）

#### 2.3.2 SKU管理接口
- CreateSKU: 创建SKU
- UpdateSKU: 更新SKU信息
- DeleteSKU: 删除SKU
- UpdateSKUStock: 更新SKU库存

#### 2.3.3 分类管理接口
- CreateCategory: 创建分类
- UpdateCategory: 更新分类信息
- DeleteCategory: 删除分类
- ListCategories: 获取分类列表
- GetCategoryTree: 获取分类树

#### 2.3.4 商品搜索接口
- SearchProducts: 搜索商品
- GetProductsByCategory: 获取分类下的商品

## 3. Proto文件结构

```proto
service ProductService {
    // 商品管理
    rpc CreateProduct(CreateProductReq) returns (CreateProductResp);
    rpc UpdateProduct(UpdateProductReq) returns (UpdateProductResp);
    rpc GetProduct(GetProductReq) returns (GetProductResp);
    rpc ListProducts(ListProductsReq) returns (ListProductsResp);
    rpc DeleteProduct(DeleteProductReq) returns (DeleteProductResp);
    rpc UpdateProductStatus(UpdateProductStatusReq) returns (UpdateProductStatusResp);
    
    // SKU管理
    rpc CreateSKU(CreateSKUReq) returns (CreateSKUResp);
    rpc UpdateSKU(UpdateSKUReq) returns (UpdateSKUResp);
    rpc DeleteSKU(DeleteSKUReq) returns (DeleteSKUResp);
    rpc UpdateSKUStock(UpdateSKUStockReq) returns (UpdateSKUStockResp);
    
    // 分类管理
    rpc CreateCategory(CreateCategoryReq) returns (CreateCategoryResp);
    rpc UpdateCategory(UpdateCategoryReq) returns (UpdateCategoryResp);
    rpc DeleteCategory(DeleteCategoryReq) returns (DeleteCategoryResp);
    rpc ListCategories(ListCategoriesReq) returns (ListCategoriesResp);
    rpc GetCategoryTree(GetCategoryTreeReq) returns (GetCategoryTreeResp);
    
    // 商品搜索
    rpc SearchProducts(SearchProductsReq) returns (SearchProductsResp);
    rpc GetProductsByCategory(GetProductsByCategoryReq) returns (GetProductsByCategoryResp);
}
```

## 4. 扩展性考虑

1. **价格系统扩展**
   - 支持多币种
   - 支持价格策略（会员价、促销价等）
   - 支持价格变更历史

2. **库存系统扩展**
   - 支持库存预占
   - 支持多仓库
   - 支持库存变更记录

3. **商品属性扩展**
   - 支持自定义属性
   - 支持属性模板
   - 支持规格组合生成SKU

4. **搜索功能扩展**
   - 支持商品标签
   - 支持商品关键词
   - 支持商品排序规则配置

## 5. Makefile更新

需要在Makefile中添加以下命令用于生成商品服务相关代码：

```makefile
.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto

.PHONY: gen-gateway-product
gen-gateway-product:
	@cd app/api && cwgo server -I ../../idl --type HTTP --service api --module ${ROOT_MOD}/app/api --idl ../../idl/api/handler_product.proto

.PHONY: gen-all-gateway
gen-all-gateway: gen-gateway-user gen-gateway-token gen-gateway-product
```

## 6. 下一步计划

1. 实现product.proto文件
2. 更新Makefile
3. 生成基础代码框架
4. 实现核心业务逻辑

是否需要我将这个设计方案写入文件？如果认可这个设计，我们可以切换到code模式来实现具体的proto文件。