# Product API Design

## Overview

This document outlines the design for the product management HTTP API. The API provides endpoints for managing products in the e-commerce system.

## API Design

### Base URL Structure
All product APIs will be prefixed with `/product`

### Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | /product/create | Create a new product |
| PUT | /product/update/:id | Update an existing product |
| DELETE | /product/:id | Delete a product |
| GET | /product/:id | Get product details |
| GET | /product/list | List products with pagination |
| GET | /product/search | Search products |

### Detailed API Specifications

#### Create Product
- Path: `/product/create`
- Method: `POST`
- Request Body:
  ```json
  {
    "name": "string",
    "description": "string",
    "picture": "string",
    "price": "float",
    "categories": ["string"]
  }
  ```
- Response:
  ```json
  {
    "base": {
      "status_code": 0,
      "status_message": "string"
    },
    "product": {
      "id": "uint32",
      "name": "string",
      "description": "string",
      "picture": "string",
      "price": "float",
      "categories": ["string"]
    }
  }
  ```

#### Update Product
- Path: `/product/update/:id`
- Method: `PUT`
- Parameters:
  - path: `id` (uint32)
- Request Body: Same as Create Product
- Response: Same as Create Product

#### Delete Product
- Path: `/product/:id`
- Method: `DELETE`
- Parameters:
  - path: `id` (uint32)
- Response:
  ```json
  {
    "base": {
      "status_code": 0,
      "status_message": "string"
    }
  }
  ```

#### Get Product
- Path: `/product/:id`
- Method: `GET`
- Parameters:
  - path: `id` (uint32)
- Response: Same as Create Product

#### List Products
- Path: `/product/list`
- Method: `GET`
- Parameters:
  - query: `page` (int32)
  - query: `page_size` (int32)
  - query: `category` (string, optional)
- Response:
  ```json
  {
    "base": {
      "status_code": 0,
      "status_message": "string"
    },
    "products": [{
      "id": "uint32",
      "name": "string",
      "description": "string",
      "picture": "string",
      "price": "float",
      "categories": ["string"]
    }],
    "total": "int32"
  }
  ```

#### Search Products
- Path: `/product/search`
- Method: `GET`
- Parameters:
  - query: `keywords` (string)
  - query: `page` (int32)
  - query: `page_size` (int32)
- Response: Same as List Products

## Data Structures

### ProductInfo
Main product information structure used across all APIs:

```protobuf
message ProductInfo {
    uint32 id = 1;
    string name = 2;
    string description = 3;
    string picture = 4;
    float price = 5;
    repeated string categories = 6;
}
```

### Common Response Structure
All APIs use a common base response structure:

```protobuf
message BaseResp {
    int32 status_code = 1;    // 0 for success, non-zero for specific error codes
    string status_message = 2; // Status description
}
```

## Proto File Design

```protobuf
syntax = "proto3";

package frontend.product;

import "api/api.proto";
import "api/common.proto";

option go_package = "basic/product";

// Product information structure
message ProductInfo {
    uint32 id = 1 [(api.go_tag)='json:"id"'];
    string name = 2 [(api.go_tag)='json:"name"'];
    string description = 3 [(api.go_tag)='json:"description"'];
    string picture = 4 [(api.go_tag)='json:"picture"'];
    float price = 5 [(api.go_tag)='json:"price"'];
    repeated string categories = 6 [(api.go_tag)='json:"categories"'];
}

// Create product
message CreateProductRequest {
    string name = 1 [(api.form)="name"];
    string description = 2 [(api.form)="description"];
    string picture = 3 [(api.form)="picture"];
    float price = 4 [(api.form)="price"];
    repeated string categories = 5 [(api.form)="categories"];
}

message CreateProductResponse {
    common.BaseResp base = 1;
    ProductInfo product = 2;
}

// Update product
message UpdateProductRequest {
    uint32 id = 1 [(api.path)="id"];
    string name = 2 [(api.form)="name"];
    string description = 3 [(api.form)="description"];
    string picture = 4 [(api.form)="picture"];
    float price = 5 [(api.form)="price"];
    repeated string categories = 6 [(api.form)="categories"];
}

message UpdateProductResponse {
    common.BaseResp base = 1;
    ProductInfo product = 2;
}

// Delete product
message DeleteProductRequest {
    uint32 id = 1 [(api.path)="id"];
}

message DeleteProductResponse {
    common.BaseResp base = 1;
}

// Get product details
message GetProductRequest {
    uint32 id = 1 [(api.path)="id"];
}

message GetProductResponse {
    common.BaseResp base = 1;
    ProductInfo product = 2;
}

// Product list
message ListProductsRequest {
    int32 page = 1 [(api.query)="page"];
    int32 page_size = 2 [(api.query)="page_size"];
    string category = 3 [(api.query)="category"];
}

message ListProductsResponse {
    common.BaseResp base = 1;
    repeated ProductInfo products = 2;
    int32 total = 3;
}

// Search products
message SearchProductsRequest {
    string keywords = 1 [(api.query)="keywords"];
    int32 page = 2 [(api.query)="page"];
    int32 page_size = 3 [(api.query)="page_size"];
}

message SearchProductsResponse {
    common.BaseResp base = 1;
    repeated ProductInfo products = 2;
    int32 total = 3;
}

// Product management service
service ProductHandler {
    // Create product
    rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse) {
        option (api.post) = "/product/create";
    }
    
    // Update product
    rpc UpdateProduct(UpdateProductRequest) returns (UpdateProductResponse) {
        option (api.put) = "/product/update/:id";
    }
    
    // Delete product
    rpc DeleteProduct(DeleteProductRequest) returns (DeleteProductResponse) {
        option (api.delete) = "/product/:id";
    }
    
    // Get product details
    rpc GetProduct(GetProductRequest) returns (GetProductResponse) {
        option (api.get) = "/product/:id";
    }
    
    // Get product list
    rpc ListProducts(ListProductsRequest) returns (ListProductsResponse) {
        option (api.get) = "/product/list";
    }
    
    // Search products
    rpc SearchProducts(SearchProductsRequest) returns (SearchProductsResponse) {
        option (api.get) = "/product/search";
    }
}
```

## Future Considerations

1. **Batch Operations**
   - Consider adding batch create/update/delete operations if needed
   - Example: POST /product/batch/create for creating multiple products

2. **Image Upload**
   - Consider adding a separate endpoint for image upload
   - Example: POST /product/upload-image

3. **Category Management**
   - Consider adding separate endpoints for category management
   - Example: CRUD operations for categories