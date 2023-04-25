# API LIST

#### 新增
```
[POST] /insert
必填 name price weight
參數 {"name":"bbbb","price":100,"weight":20,"tiket":0,"freight":0,"fare":0,"status":0, "exchange_rate": 0.23,"profit":10}
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 100001
}
```

#### 查詢
```
[GET] /quote?name=&offset=1
參數 name offset
```

#### Response
```bash
{
    "data": [
        {
            "id": 11,
            "name": "aaa",
            "price": 100,
            "weight": 20,
            "exchange_rate": 0.23,
            "profit": 10,
            "created_at": "2023-04-25T10:15:42Z",
            "updated_at": "2023-04-25T10:15:42Z"
        },
        {
            "id": 12,
            "name": "aaa",
            "price": 100,
            "weight": 20,
            "exchange_rate": 0.23,
            "profit": 10,
            "created_at": "2023-04-25T10:15:42Z",
            "updated_at": "2023-04-25T10:15:42Z"
        },
        {
            "id": 14,
            "name": "bbbb",
            "price": 100,
            "weight": 20,
            "exchange_rate": 0.23,
            "profit": 10,
            "created_at": "2023-04-25T10:15:46Z",
            "updated_at": "2023-04-25T10:15:46Z"
        }
    ]
}


{
     "error_code": 100001
}
```

#### 刪除
```
[DELETE] /:id
參數 id
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 100001
}
```

#### 更新
```
[PUT] /update
必填 id
參數 {"id": 1,"name":"bbbb","price":100,"weight":20,"tiket":0,"freight":0,"fare":0,"status":0, "exchange_rate": 0.23,"profit":10}
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 100001
}
```