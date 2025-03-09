# Product Service API Test Cases

This document provides curl commands for testing the product service HTTP endpoints.

## Create Product
```bash
curl -X POST http://localhost:8080/product/manage \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "name=Test Product" \
  -d "description=A test product description" \
  -d "picture=test.jpg" \
  -d "price=99.99" \
  -d "categories=Electronics&categories=Gadgets"
```

Expected Success Response:
```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  },
  "product": {
    "id": 1,
    "name": "Test Product",
    "description": "A test product description",
    "picture": "test.jpg",
    "price": 99.99,
    "categories": ["Electronics", "Gadgets"]
  }
}
```

## Get Product
```bash
curl -X GET http://localhost:8080/product/detail/1
```

Expected Success Response:
```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  },
  "product": {
    "id": 1,
    "name": "Test Product",
    "description": "A test product description",
    "picture": "test.jpg",
    "price": 99.99,
    "categories": ["Electronics", "Gadgets"]
  }
}
```

## Update Product
```bash
curl -X PUT http://localhost:8080/product/manage/update/1 \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "name=Updated Product" \
  -d "description=An updated product description" \
  -d "picture=updated.jpg" \
  -d "price=199.99" \
  -d "categories=Electronics&categories=Premium"
```

Expected Success Response:
```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  },
  "product": {
    "id": 1,
    "name": "Updated Product",
    "description": "An updated product description",
    "picture": "updated.jpg",
    "price": 199.99,
    "categories": ["Electronics", "Premium"]
  }
}
```

## Delete Product
```bash
curl -X DELETE http://localhost:8080/product/manage/delete/1
```

Expected Success Response:
```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  }
}
```

## List Products
```bash
# List all products
curl -X GET "http://localhost:8080/products?page=1&page_size=10"

# List products by category
curl -X GET "http://localhost:8080/products?page=1&page_size=10&category=Electronics"
```

Expected Success Response:
```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  },
  "products": [
    {
      "id": 1,
      "name": "Test Product",
      "description": "A test product description",
      "picture": "test.jpg",
      "price": 99.99,
      "categories": ["Electronics", "Gadgets"]
    }
  ],
  "total": 1
}
```

## Search Products
```bash
curl -X GET "http://localhost:8080/products/search?keywords=test&page=1&page_size=10"
```

Expected Success Response:
```json
{
  "base": {
    "status_code": 0,
    "status_message": "success"
  },
  "products": [
    {
      "id": 1,
      "name": "Test Product",
      "description": "A test product description",
      "picture": "test.jpg",
      "price": 99.99,
      "categories": ["Electronics", "Gadgets"]
    }
  ],
  "total": 1
}
```

## Error Response Example
All endpoints return the following format for errors:
```json
{
  "base": {
    "status_code": 500,
    "status_message": "Error message details"
  }
}