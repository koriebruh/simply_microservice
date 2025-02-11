### *`CREATE ORDER`*

```
POST /api/orders
Content-Type: application/json

REQUEST =>
{
    "customer_id" : "1",
    "items" : [
      {"product_id" : 10, "quantity" : 2},  
      {"product_id" : 13, "quantity" : 1},  
    ],
    "shiping_addr" : "jalan baru jadi",
    "payment_method" : "BCA Transfer",
    "ammount" : 100000
}   

RESPONSE =>
{
    "status" : "success"
    "message" : "create order success"
}
```

### *`STATUS ORDER`*

```
GET /api/orders/status
Content-Type: application/json

RESPONSE =>
{
    "status" : "success"
    "message" : "status order"
    "data" : {
        "items" : [
          {"product_id" : 10, "quantity" : 2, price_per_item : 10000},  
          {"product_id" : 13, "quantity" : 1, price_per_item : 10000},  
        ],
        "ammount" : 100000,
        "payment_method" : "BCA Transfer",
        "shiping_addr" : "jalan baru jadi",
        "shipping status" : "packing"
    }
}